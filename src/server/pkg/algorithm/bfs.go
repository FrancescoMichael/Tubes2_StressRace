package algorithm

import (
	scraper "server/pkg/scraper"
	"strings"
)

func Bfs(start string, end string) []string {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string]string)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return makePath(parent, start, end)
		}

		var allUrl = scraper.GetScrapeLinksColly(curr)
		if allUrl == nil {
			continue // Handle nil or handle error if function can error out
		}

		for _, linkTemp := range allUrl {
			if !visited[linkTemp] {
				visited[linkTemp] = true
				parent[linkTemp] = curr
				queue = append(queue, linkTemp)
				if linkTemp == end {
					return makePath(parent, start, end)
				}
			}
		}
	}

	return nil // Return nil if end is not reachable
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
