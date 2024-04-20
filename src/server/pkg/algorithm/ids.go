package algorithm

import (
	scraper "server/pkg/scraper"
	"strings"
	"sync"
)

// find all paths version
// var mutex sync.Mutex

func Ids(startPage string, endPage string, maxDepth int) ([][]string, int) {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)
	var paths [][]string       // store possible paths from start to goal page
	var shortestDepth int = -1 // depth of shortest path

	for depth := 0; depth <= maxDepth && shortestDepth == -1; depth++ {
		var currPathsDepth [][]string                       // store possible paths temporary
		var visited map[string]bool = make(map[string]bool) // store nodes/links that has been visited to avoid cycles
		dfs(startPage, endPage, depth, visited, nil, &currPathsDepth)
		// if possible path has bee n found
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
	var allUrl = scraper.GetScrapeLinksColly(currUrl)
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

	var allUrl = scraper.GetScrapeLinksColly(currUrl)
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

var mutex sync.Mutex
var foundGlobal bool // global indicator if path is found

func IdsFirstGoRoutine(startPage string, endPage string, maxDepth int) []string {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)
	var result []string
	var wg sync.WaitGroup

	for depth := 0; depth < maxDepth && !foundGlobal; depth++ {
		visited := make(map[string]bool)
		path := []string{}
		wg.Add(1)

		go func(d int) {
			defer wg.Done()
			if found, res := DfsFirstGoRoutine(startPage, endPage, d, visited, path); found {
				mutex.Lock()
				if !foundGlobal { // Ensure no other goroutine has set the path
					result = res
					foundGlobal = true
				}
				mutex.Unlock()
			}
		}(depth)
		wg.Wait() // Wait for the goroutines of this depth level
		if foundGlobal {
			break
		}
	}

	return result
}

func DfsFirstGoRoutine(currUrl string, endPage string, depth int, visited map[string]bool, path []string) (bool, []string) {
	if foundGlobal {
		return false, nil // Stop processing if the path is already found
	}

	if currUrl == endPage {
		return true, append(path, currUrl)
	}
	if depth <= 0 {
		return false, nil
	}

	mutex.Lock() // Lock before accessing global resources
	if visited[currUrl] {
		mutex.Unlock()
		return false, nil
	}
	visited[currUrl] = true
	mutex.Unlock()

	path = append(path, currUrl)
	var result []string
	var found bool

	allUrl := scraper.GetScrapeLinksColly(currUrl)
	if allUrl == nil {
		return false, nil
	}

	for _, value := range allUrl {
		if !visited[value] {
			if f, res := DfsFirst(value, endPage, depth-1, visited, path); f {
				result = res
				found = f
				break
			}
		}
	}

	mutex.Lock()
	visited[currUrl] = false // Unmark the current node after finishing
	mutex.Unlock()

	return found, result
}
