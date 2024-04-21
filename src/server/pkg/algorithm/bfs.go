package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
)

func Bfs(start string, end string) ([]string, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(start) {
		return nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(end) {
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
