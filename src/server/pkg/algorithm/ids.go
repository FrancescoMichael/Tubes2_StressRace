package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
)

var LockPaths sync.Mutex
func IdsFirstPath(startUrl string, endUrl string, maxDepth int) ([]string, map[string]bool, error) {
	startUrl = strings.TrimSpace(startUrl)
	endUrl = strings.TrimSpace(endUrl)
	var resultPath []string
	if !scraper.IsWikiPageUrlExists(&startUrl) {
		return nil, nil, fmt.Errorf("start page doesn't exist")
	}
	if !scraper.IsWikiPageUrlExists(&endUrl) {
		return nil, nil, fmt.Errorf("end page doesn't exist")
	}

	for depth := 0; depth <= maxDepth; depth++ {
		dlsFirstPath(startUrl, endUrl, depth, &resultPath, []string{})
		if len(resultPath) > 0 {
			return resultPath, nil, nil
		}
		// fmt.Println("HERE")
	}
	return nil, nil, nil
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

func IdsAllPath(startUrl string, endUrl string, maxDepth int) ([][]string, map[string]bool, error) {
	startUrl = strings.TrimSpace(startUrl)
	endUrl = strings.TrimSpace(endUrl)
	var allPaths [][]string
	if !scraper.IsWikiPageUrlExists(&startUrl) {
		return nil, nil, fmt.Errorf("start page doesn't exist")
	}
	if !scraper.IsWikiPageUrlExists(&endUrl) {
		return nil, nil, fmt.Errorf("end page doesn't exist")
	}

	for depth := 0; depth <= maxDepth; depth++ {
		dlsAllPath(startUrl, endUrl, depth, &allPaths, []string{})
		if len(allPaths) > 0 {
			return allPaths, nil, nil
		}
		// fmt.Println("HERE")
	}
	return nil, nil, nil
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
