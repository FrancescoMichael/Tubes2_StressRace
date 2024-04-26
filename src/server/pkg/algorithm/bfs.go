package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
	"sync/atomic"
)

// Basic node structure that BfsGoRoutine and BfsGoRoutineAllPaths
type node struct {
	url   string
	depth int32
}

// basic function to append node
func appendNode(list []node, newNode node) []node {
	return append(list, newNode)
}

// this is the most simple version of BFS, all BFS versions will be based on this
func Bfs(start string, end string) ([]string, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) { // Check if start url exists
		return nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) { // Check if end/targer url exists
		return nil, fmt.Errorf("end page does not exist")
	}
	queue := []string{start}          // FIFO for BFS
	visited := make(map[string]bool)  // Visited map to avoid cycles
	visited[start] = true             // mark the first as true
	parent := make(map[string]string) // Make parent map to keep track of parent nodes
	var mutex1Queue sync.Mutex        // mutex1 used when making path, technically it is not used in this BFS func but need in the makePath function

	for len(queue) > 0 { // iterate until queue is empty or path is found
		curr := queue[0]  // curr node , using FIFO
		queue = queue[1:] // dequeue

		if curr == end { // return path if curr url is target
			return makePath(parent, start, end, &mutex1Queue), nil
		}

		var allUrl = scraper.GetScrapeLinks(curr) // find all child nodes
		if allUrl == nil {
			continue // skip if there isn't any child
		}

		for _, linkTemp := range allUrl { // iterate over every child
			if !visited[linkTemp] {
				visited[linkTemp] = true
				parent[linkTemp] = curr
				queue = append(queue, linkTemp) // enqueue
				if linkTemp == end {            // check if node is target, avoid unecessary process later on
					return makePath(parent, start, end, &mutex1Queue), nil
				}
			}
		}
	}

	return nil, nil // Return nil if end is not reachable
}

// BfsGoRoutine is the same as Bfs, but it uses go routine
func BfsGoRoutine(start string, end string) ([]string, map[string]bool, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, nil, fmt.Errorf("end page does not exist")
	}
	limiter := make(chan struct{}, 100) // set limiter of go routines
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string]string)
	var found int32 = 0 // Use atomic int32 for the found flag
	var mutex1Queue sync.Mutex

	for atomic.LoadInt32(&found) == 0 { // iterate until found a solution
		limiter <- struct{}{} // insert limiter

		go func() {
			defer func() { <-limiter }() // release limiter when process is finished
			var curr string

			mutex1Queue.Lock()
			if len(queue) > 0 {
				curr = queue[0]
				queue = queue[1:]
			}
			mutex1Queue.Unlock()

			if curr == end { // Set 'found' to 1 if a solution is discovered.
				atomic.StoreInt32(&found, 1)
				return
			}

			var allUrl = scraper.GetScrapeLinksConcurrent(curr)
			if allUrl == nil {
				return
			}

			mutex1Queue.Lock()
			for _, linkTemp := range allUrl {

				if !visited[linkTemp] {
					visited[linkTemp] = true
					parent[linkTemp] = curr
					queue = append(queue, linkTemp)
					if linkTemp == end {
						atomic.StoreInt32(&found, 1) // Set 'found' to 1 if a solution is discovered.
						break
					}
				}
			}
			mutex1Queue.Unlock()

		}()
	}

	if atomic.LoadInt32(&found) == 0 {
		return nil, visited, fmt.Errorf("path cannot be found") // Return nil if end is not reachable
	}

	return makePath(parent, start, end, &mutex1Queue), visited, nil
}

// this is the same as BFS but uses node to keep track of current depth
// it will search for all shortest path in the shortest depth possible
func BfsMultPath(start string, end string) ([][]string, map[string]bool, error) {
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)

	if !scraper.IsWikiPageUrlExists(&start) {
		return nil, nil, fmt.Errorf("start page does not exist")
	}
	if !scraper.IsWikiPageUrlExists(&end) {
		return nil, nil, fmt.Errorf("end page does not exist")
	}
	startNode := node{ // initialize startNode
		url:   start,
		depth: 0,
	}
	var queue []node
	queue = appendNode(queue, startNode)
	visited := make(map[string]bool)
	visited[start] = true
	parent := make(map[string][]string)
	// parents := make(map[string][]string)
	var maxDepth int32 = 999 // records the depth of the shortest solution

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.depth > maxDepth { // stops when the current depth is bigger then the shortest depth solution
			break
		}

		if curr.url == end && curr.depth <= maxDepth {
			maxDepth = curr.depth
			if curr.depth == 0 || curr.depth == 1 {
				break
			}

		}

		var allUrl = scraper.GetScrapeLinks(curr.url)
		if allUrl == nil {
			continue // continue when current node is a leaf / no children
		}
		currDepth := curr.depth
		for _, linkTemp := range allUrl {
			if !visited[linkTemp] {
				visited[linkTemp] = true
				parent[linkTemp] = append(parent[linkTemp], curr.url)
				newNode := node{ // initialize node with depth + 1 of parent node
					url:   linkTemp,
					depth: currDepth + 1,
				}
				queue = appendNode(queue, newNode) // enqueue

			}
		}
	}

	return makePathAll(parent, start, end), visited, nil
}

// func BfsAllPathGoRoutine is the same as BfsAllPath but it uses go routine to speed up process
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
		wg.Add(1) // using wait group to make sure all go routines has finished
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

			if curr.url == end && curr.depth <= atomic.LoadInt32(&maxDepth) { // when current url is the target, update maxDepth
				atomic.SwapInt32(&maxDepth, curr.depth)
				return
			}
			if curr.depth >= atomic.LoadInt32(&maxDepth) { // when current depth is more than the maxDepth, make found = 1 as a flag to not process nodes anymore
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
					parent[linkTemp] = append(parent[linkTemp], curr.url) // update parent map
					if linkTemp == end && curr.depth <= atomic.LoadInt32(&maxDepth) {
						atomic.SwapInt32(&maxDepth, curr.depth+1)
						visited[linkTemp] = false
						break
					} else {
						queue = append(queue, node{ // enqueue node, node has depth + 1 of parent
							url:   linkTemp,
							depth: curr.depth + 1,
						})

					}

				}
			}
			mutex1.Unlock()

		}()
	}
	wg.Wait() // waits until all go routines are finished

	if atomic.LoadInt32(&found) == 0 {
		return nil, visited, fmt.Errorf("path cannot be found") // Return nil if end is not reachable
	}

	return makePathAll(parent, start, end), visited, nil
}

// func makePath receives parent map, start url, and end url to create a path from end to start based on parent map
func makePath(parent map[string]string, start string, end string, mutexQueue *sync.Mutex) []string {
	path := []string{}
	curr := end
	for curr != start { // iterate until curr is start url
		path = append([]string{curr}, path...)
		mutexQueue.Lock()
		curr = parent[curr]
		mutexQueue.Unlock()

	}
	path = append([]string{start}, path...)
	return path
}

// func makePath receives parent map, start url, and end url to create all paths from start url to end url
func makePathAll(parent map[string][]string, start string, end string) [][]string {
	if start == end { // base case, start is end
		return [][]string{{start}}
	}
	var paths [][]string // initialize paths

	if _, exists := parent[end]; !exists { // check if end has a parent
		return nil
	}

	for _, parentUrl := range parent[end] {

		parentPaths := makePathAll(parent, start, parentUrl) //recursively call itself to find all paths from start url to parent url

		for _, path := range parentPaths {
			paths = append(paths, append(path, end)) // append every parent paths with end url and insert it as new paths
		}
	}

	return paths
}
