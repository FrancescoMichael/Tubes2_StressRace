package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
)

var LockPaths sync.Mutex

// go routine version of IDS will be based on this Ids
func Ids(startPage string, endPage string, maxDepth int) ([][]string, int) {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)
	var paths [][]string       // store possible paths from start to goal page
	var shortestDepth int = -1 // depth of shortest path

	for depth := 0; depth <= maxDepth && shortestDepth == -1; depth++ {
		var currPathsDepth [][]string                       // store possible paths temporary
		var visited map[string]bool = make(map[string]bool) // store nodes/links that has been visited to avoid cycles
		dls(startPage, endPage, depth, visited, nil, &currPathsDepth)
		// if possible path has bee n found
		if len(currPathsDepth) > 0 {
			shortestDepth = depth
			paths = append(paths, currPathsDepth...)
		}

	}

	return paths, shortestDepth
}

func dls(currUrl string, endPage string, currDepth int, visited map[string]bool, currPath []string, currPathDepth *[][]string) {
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
			dls(url, endPage, currDepth-1, visited, currPath, currPathDepth)
		}
	}
	visited[currUrl] = false

}

func IdsFirstPath(startUrl string, endUrl string, maxDepth int) ([]string, int, error) {
	startUrl = strings.TrimSpace(startUrl)
	endUrl = strings.TrimSpace(endUrl)
	var resultPath []string
	if !scraper.IsWikiPageUrlExists(&startUrl) {
		return nil, 0, fmt.Errorf("start page doesn't exist")
	}
	if !scraper.IsWikiPageUrlExists(&endUrl) {
		return nil, 0, fmt.Errorf("end page doesn't exist")
	}

	for depth := 0; depth <= maxDepth; depth++ {
		dlsFirstPath(startUrl, endUrl, depth, &resultPath, []string{})
		if len(resultPath) > 0 {
			return resultPath, len(scraper.Unique), nil
		}
		// fmt.Println("HERE")
	}
	return nil, len(scraper.Unique), fmt.Errorf("path cannot be found at depth 10")
}

func dlsFirstPath(currUrl string, endUrl string, depth int, resultPaths *[]string, currPath []string) {
	// fmt.Println(currUrl)
	LockPaths.Lock()
	newPath := append(currPath, currUrl)
	if currUrl == endUrl {
		*resultPaths = append(*resultPaths, newPath...)
		LockPaths.Unlock()
		return
	} else if len(*resultPaths) > 0 {
		LockPaths.Unlock()
		return
	}
	LockPaths.Unlock()

	if depth <= 0 {
		return
	}

	limiter := make(chan struct{}, 15)
	for _, urlNow := range scraper.GetScrapeLinksConcurrent(currUrl) {
		limiter <- struct{}{}
		go func(nextUrl string) {
			dlsFirstPath(nextUrl, endUrl, depth-1, resultPaths, newPath)
			<-limiter
		}(urlNow)
	}
}

func IdsAllPath(startUrl string, endUrl string, maxDepth int) ([][]string, int, error) {
	startUrl = strings.TrimSpace(startUrl)
	endUrl = strings.TrimSpace(endUrl)
	var allPaths [][]string
	if !scraper.IsWikiPageUrlExists(&startUrl) {
		return nil, 0, fmt.Errorf("start page doesn't exist")
	}
	if !scraper.IsWikiPageUrlExists(&endUrl) {
		return nil, 0, fmt.Errorf("end page doesn't exist")
	}

	for depth := 0; depth <= maxDepth; depth++ {
		dlsAllPath(startUrl, endUrl, depth, &allPaths, []string{})
		if len(allPaths) > 0 {
			return allPaths, len(scraper.Unique), nil
		}
		// fmt.Println("HERE")
	}
	return nil, len(scraper.Unique), nil
}
func dlsAllPath(currUrl string, endUrl string, depth int, paths *[][]string, currPath []string) {

	if currUrl == endUrl {
		// fmt.Println(currUrl)
		LockPaths.Lock()
		newPath := append(currPath, currUrl)
		*paths = append(*paths, newPath)
		LockPaths.Unlock()
		return
	}

	if depth <= 0 {
		return
	}
	newPath := append(currPath, currUrl)

	limiter := make(chan struct{}, 15)
	for _, urlNow := range scraper.GetScrapeLinksConcurrent(currUrl) {
		limiter <- struct{}{}
		go func(nextUrl string) {
			dlsAllPath(nextUrl, endUrl, depth-1, paths, newPath)
			<-limiter
		}(urlNow)
	}
}
