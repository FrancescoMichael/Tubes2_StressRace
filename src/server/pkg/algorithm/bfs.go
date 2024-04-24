package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
	"sync/atomic"
)

type node struct {
	url   string
	depth int32
}

func appendNode(list []node, newNode node) []node {
	return append(list, newNode)
}

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
	// parents := make(map[string][]string)

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

func BfsGoRoutine(start string, end string) ([]string, map[string]bool, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, nil, fmt.Errorf("end page does not exist")
	}
	limiter := make(chan struct{}, 100)
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string]string)
	var found int32 = 0 // Use atomic int32 for the found flag
	var mutex1 sync.Mutex

	for atomic.LoadInt32(&found) == 0 {
		limiter <- struct{}{}

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
		return nil, visited, fmt.Errorf("path cannot be found") // Return nil if end is not reachable
	}

	return makePath(parent, start, end), visited, nil
}

func BfsMultPath(start string, end string) ([][]string, map[string]bool, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, nil, fmt.Errorf("end page does not exist")
	}
	startNode := node{
		url:   start,
		depth: 0,
	}
	var queue []node
	queue = appendNode(queue, startNode)
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string][]string)
	// parents := make(map[string][]string)
	var maxDepth int32 = 999

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.depth > maxDepth {
			break
		}

		if curr.url == end && curr.depth <= maxDepth {
			// fmt.Println(curr.url)
			maxDepth = curr.depth
			// fmt.Println(len(queue))
			if curr.depth == 0 || curr.depth == 1 {
				break
			}

		}

		var allUrl = scraper.GetScrapeLinks(curr.url)
		if allUrl == nil {
			continue // Handle nil or handle error if function can error out
		}
		currDepth := curr.depth
		// fmt.Println(currDepth)
		for _, linkTemp := range allUrl {
			if !visited[linkTemp] {
				visited[linkTemp] = true
				parent[linkTemp] = append(parent[linkTemp], curr.url)
				newNode := node{
					url:   linkTemp,
					depth: currDepth + 1,
				}
				queue = appendNode(queue, newNode)

			}
		}
	}

	return makePathAll(parent, start, end), visited, nil // Return nil if end is not reachable
}

func BfsAllPathGoRoutine(start string, end string) ([][]string, map[string]bool, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, nil, fmt.Errorf("end page does not exist")
	}
	limiter := make(chan struct{}, 100)
	startNode := node{
		url:   start,
		depth: 0,
	}
	queue := []node{startNode}

	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string][]string)
	var found int32 = 0 // Use atomic int32 for the found flag
	var mutex1 sync.Mutex
	var wg sync.WaitGroup
	var maxDepth int32 = 10

	for atomic.LoadInt32(&found) == 0 {
		limiter <- struct{}{}
		wg.Add(1)
		go func() {

			defer func() { <-limiter }()
			defer wg.Done()
			var curr node

			mutex1.Lock()
			if len(queue) > 0 {
				curr = queue[0]
				queue = queue[1:]
			}
			mutex1.Unlock()

			if curr.url == end && curr.depth <= atomic.LoadInt32(&maxDepth) {
				atomic.SwapInt32(&maxDepth, curr.depth)
				return
			}
			if curr.depth >= atomic.LoadInt32(&maxDepth) {
				atomic.StoreInt32(&found, 1)
				return
			}

			var allUrl = scraper.GetScrapeLinksConcurrent(curr.url)
			if allUrl == nil {
				return
			}

			mutex1.Lock()
			for _, linkTemp := range allUrl {

				if !visited[linkTemp] {
					visited[linkTemp] = true
					parent[linkTemp] = append(parent[linkTemp], curr.url)
					if linkTemp == end && curr.depth <= atomic.LoadInt32(&maxDepth) {
						atomic.SwapInt32(&maxDepth, curr.depth+1)
						visited[linkTemp] = false
						break
					} else {
						queue = append(queue, node{
							url:   linkTemp,
							depth: curr.depth + 1,
						})

					}

				}
			}
			mutex1.Unlock()

		}()
	}
	wg.Wait()

	if atomic.LoadInt32(&found) == 0 {
		return nil, visited, fmt.Errorf("path cannot be found") // Return nil if end is not reachable
	}

	return makePathAll(parent, start, end), visited, nil
}

// func BfsMultPathGoRoutine(start string, end string) ([][]string, map[string]bool, error) {
// 	start = strings.TrimSpace(start)
// 	end = strings.TrimSpace(end)

// 	if !scraper.IsWikiPageUrlExists(&start) {
// 		return nil, nil, fmt.Errorf("start page does not exist")
// 	}
// 	if !scraper.IsWikiPageUrlExists(&end) {
// 		return nil, nil, fmt.Errorf("end page does not exist")
// 	}
// 	startNode := node{
// 		url:   start,
// 		depth: 0,
// 	}
// 	var queue []node
// 	queue = appendNode(queue, startNode)
// 	visited := make(map[string]bool)
// 	visited[start] = true
// 	parent := make(map[string][]string)
// 	// parents := make(map[string][]string)
// 	var maxDepth = 10
// 	limiter := make(chan struct{}, 100)
// 	var mutex sync.Mutex
// 	var wg sync.WaitGroup

// 	for {
// 		var curr node
// 		mutex.Lock()
// 		if len(queue) > 0 {
// 			curr = queue[0]
// 			queue = queue[1:]
// 			// fmt.Println(len(queue))

// 		}

// 		mutex.Unlock()

// 		if curr.depth > maxDepth {
// 			fmt.Println("STUCK HERE")
// 			break
// 		}

// 		if curr.url == end && curr.depth <= maxDepth {
// 			mutex.Lock()
// 			fmt.Println("ATAS")
// 			maxDepth = curr.depth
// 			mutex.Unlock()
// 			continue
// 		}
// 		if curr.url == "https://en.wikipedia.org/wiki/Pancasila_(politics)" || curr.url == "https://en.wikipedia.org/wiki/Slovakia" {
// 			fmt.Println(curr.url)
// 		}
// 		wg.Add(1)
// 		limiter <- struct{}{}
// 		go func(curr node) {
// 			defer wg.Done()
// 			defer func() { <-limiter }()
// 			allUrl := scraper.GetScrapeLinksConcurrent(curr.url)
// 			if allUrl == nil {
// 				return
// 			}

// 			mutex.Lock()
// 			for _, linkTemp := range allUrl {
// 				if !visited[linkTemp] {
// 					// if curr.url == "https://en.wikipedia.org/wiki/Pancasila_(politics)" || curr.url == "https://en.wikipedia.org/wiki/Slovakia" {
// 					// 	// fmt.Println(curr.url)
// 					// 	// fmt.Println(linkTemp)

// 					// }
// 					visited[linkTemp] = true
// 					parent[linkTemp] = append(parent[linkTemp], curr.url)
// 					newNode := node{
// 						url:   linkTemp,
// 						depth: curr.depth + 1,
// 					}
// 					if linkTemp == end && curr.depth+1 <= maxDepth {

// 						fmt.Println("BAWAH")
// 						maxDepth = curr.depth + 1
// 						return

// 					} else {
// 						queue = appendNode(queue, newNode)
// 					}

// 				}
// 			}
// 			mutex.Unlock()
// 		}(curr)
// 	}

// 	wg.Wait() // Wait for all go routines to finish

// 	return makePathAll(parent, start, end), visited, nil // Return nil if end is not reachable
// }

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
func makePathAll(parent map[string][]string, start string, end string) [][]string {
	if start == end {
		return [][]string{{start}}
	}
	var paths [][]string

	if _, exists := parent[end]; !exists {
		return nil
	}

	for _, p := range parent[end] {

		parentPaths := makePathAll(parent, start, p)

		for _, path := range parentPaths {
			paths = append(paths, append(path, end))
		}
	}

	return paths
}
