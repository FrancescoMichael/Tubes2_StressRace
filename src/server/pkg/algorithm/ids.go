package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
)

var LockPaths sync.Mutex

// go routine version of IDS will be based on this Ids
// func Ids receives startPage and endPage URL, and returns the list of all paths using IDS algorithm
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

// func dls will peform dfs with a depth constraint
// dls will peform dfs recursively
func dls(currUrl string, endPage string, currDepth int, visited map[string]bool, currPath []string, currPathDepth *[][]string) {
	if currDepth == 0 && currUrl == endPage { // if the currUrl is the target Url, then update currentPathDepth
		path := make([]string, len(currPath)+1)
		copy(path, append(currPath, currUrl))
		*currPathDepth = append(*currPathDepth, path)
		return

	}

	if currDepth <= 0 { //  if current depth is at the bottom of dfs, then stop recursion
		return
	}

	var allUrl = scraper.GetScrapeLinks(currUrl) // getting all possible child of currUrl
	if allUrl == nil {                           // if no more child (url is a leaf) then return
		return
	}
	visited[currUrl] = true              // mark current node/url as already visited
	currPath = append(currPath, currUrl) // update current path

	for _, url := range allUrl { // iterate over every child
		if !visited[url] {
			dls(url, endPage, currDepth-1, visited, currPath, currPathDepth) // process left most child first
		}
	}
	visited[currUrl] = false // backtrack

}

// func IdsFirstPath is IDS with go routine implementation and only gives one solution
// in func IdsFirstPath, there is no visited map to keep track of all nodes that has been visited
// because DLS basically DFS but with a maxDepth, therefore Infinite Loops in Cycles is impossible
// reference : https://youtu.be/Y85ECk_H3h4?si=dm3PtHyv16rDHC38
// The reason that IdsFirstPath does not have visited map is to reduce the amount of mutex used when using go routine
// IdsFirstPath is the same as IDS, but it only returns one path and it also returns the amount of articles that has been processed
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
	}
	return nil, len(scraper.Unique), fmt.Errorf("path cannot be found at depth 10")
}

// dlsFirstPath functions similarly to dls, but it utilizes goroutines and terminates as soon as a path is found
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

	limiter := make(chan struct{}, 15) // limits how many go routines are running concurrently
	for _, urlNow := range scraper.GetScrapeLinksConcurrent(currUrl) {
		limiter <- struct{}{}
		go func(nextUrl string) {
			dlsFirstPath(nextUrl, endUrl, depth-1, resultPaths, newPath)
			<-limiter
		}(urlNow)
	}
}

// Func IdsAllPath is the same as Ids, but it uses go routines
// similar to IdsFirstPath, IdsAlPathh doesn't have visited map to reduce mutex
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
	}
	return nil, len(scraper.Unique), nil
}

// Func dlsAllPath is the same as dls, but it uses go routines
// similar to IdsFirstPath, IdsAlPathh doesn't have visited map to reduce mutex
func dlsAllPath(currUrl string, endUrl string, depth int, paths *[][]string, currPath []string) {

	if currUrl == endUrl {
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
