package algorithm

import (
	scraper "server/pkg/scraper"
	"strings"
)

// find all paths version
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
	var allUrl = scraper.GetScrapeLinks(currUrl)
	if allUrl == nil {
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

// find path and then exit version

func IdsFirst(startPage string, endPage string, maxDepth int) []string {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)
	for depth := 0; depth < maxDepth; depth++ {
		visited := make(map[string]bool)
		path := []string{}
		if found, result := DfsFirst(startPage, endPage, depth, visited, path); found {
			return result
		}

	}
	return nil
}

func DfsFirst(currUrl string, endPage string, depth int, visited map[string]bool, path []string) (bool, []string) {
	if currUrl == endPage {
		return true, append(path, currUrl)
	}
	if depth <= 0 {
		return false, nil
	}

	var allUrl = scraper.GetScrapeLinks(currUrl)
	if allUrl == nil {
		return false, nil
	}

	visited[currUrl] = true
	path = append(path, currUrl)

	for _, value := range allUrl {
		if !visited[value] {
			if found, result := DfsFirst(value, endPage, depth-1, visited, path); found {
				return true, result
			}
		}
	}

	visited[currUrl] = false // Unmark the current node
	return false, nil
}
