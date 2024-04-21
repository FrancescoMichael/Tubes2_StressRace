package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
	"sync/atomic"
)

func Bfs(start string, end string) ([]string, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, fmt.Errorf("end page does not exist")
	}
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string]string)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return makePath(parent, start, end), nil
		}

		var allUrl = scraper.GetScrapeLinks(curr)
		if allUrl == nil {
			continue // Handle nil or handle error if function can error out
		}

		for _, linkTemp := range allUrl {
			if !visited[linkTemp] {
				visited[linkTemp] = true
				parent[linkTemp] = curr
				queue = append(queue, linkTemp)
				if linkTemp == end {
					return makePath(parent, start, end), nil
				}
			}
		}
	}

	return nil, nil // Return nil if end is not reachable
}

func BfsGoRoutine(start string, end string) ([]string, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, fmt.Errorf("end page does not exist")
	}
	limiter := make(chan int, 50)
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string]string)
	var found int32 = 0 // Use atomic int32 for the found flag
	var mutex1 sync.Mutex

	for atomic.LoadInt32(&found) == 0 {
		limiter <- 1

		go func() {
			defer func() { <-limiter }()
			var curr string

			mutex1.Lock()
			if len(queue) > 0 {
				curr = queue[0]
				queue = queue[1:]
			}
			mutex1.Unlock()

			if curr == end {
				atomic.StoreInt32(&found, 1)
				return
			}

			var allUrl = scraper.GetScrapeLinksConcurrent(curr)
			if allUrl == nil {
				return
			}

			mutex1.Lock()
			for _, linkTemp := range allUrl {
				if !visited[linkTemp] {
					visited[linkTemp] = true
					parent[linkTemp] = curr
					queue = append(queue, linkTemp)
					if linkTemp == end {
						atomic.StoreInt32(&found, 1)
						break
					}
				}
			}
			mutex1.Unlock()

		}()
	}

	if atomic.LoadInt32(&found) == 0 {
		return nil, nil // Return nil if end is not reachable
	}

	return makePath(parent, start, end), nil
}

func makePath(parent map[string]string, start string, end string) []string {
	path := []string{}
	curr := end
	for curr != start {
		path = append([]string{curr}, path...)
		curr = parent[curr]

	}
	path = append([]string{start}, path...)
	return path
}
