package algorithm

import (
	scraper "server/pkg/scraper"
	"strings"
)

func Ids(startPage string, endPage string, maxDepth int) ([][]string, int) {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)
	var paths [][]string       // store possible paths from start to goal page
	var shortestDepth int = -1 // depth of shortest path

	for depth := 0; depth <= maxDepth && shortestDepth == -1; depth++ {
		var currPathsDepth [][]string                       // store possible paths temporary
		var visited map[string]bool = make(map[string]bool) // store nodes/links that has been visited to avoid cycles
		dfs(startPage, endPage, depth, visited, nil, &currPathsDepth)
		// if possible path has been found
		if len(currPathsDepth) > 0 {
			shortestDepth = depth
			paths = append(paths, currPathsDepth...)
		}

	}

	return paths, shortestDepth
}

func dfs(currUrl string, endPage string, currDepth int, visited map[string]bool, currPath []string, currPathDepth *[][]string) {
	if currDepth == 0 && currUrl == endPage {
		path := make([]string, len(currPath)+1)
		copy(path, append(currPath, currUrl))
		*currPathDepth = append(*currPathDepth, path)

	}
	if currDepth <= 0 {
		return
	}
	var allUrl []string
	err := scraper.WebScraping(currUrl, &allUrl)
	if err != nil {
		return
	}
	visited[currUrl] = true
	currPath = append(currPath, currUrl)

	for _, url := range allUrl {
		if !visited[url] {
			dfs(url, endPage, currDepth-1, visited, currPath, currPathDepth)
		}
	}
	visited[currUrl] = false

}
