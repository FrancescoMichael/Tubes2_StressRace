package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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

func IdsFirst(startPage string, endPage string, maxDepth int) ([]string, error) {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)

	if !scraper.IsWikiPageUrlExists(&startPage) {
		return nil, fmt.Errorf("start page doesn't exists")
	}
	if !scraper.IsWikiPageUrlExists(&endPage) {
		return nil, fmt.Errorf("end page doesn't exists")
	}

	for depth := 0; depth < maxDepth; depth++ {
		visited := make(map[string]bool)
		path := []string{}
		if found, result := DlsFirst(startPage, endPage, depth, visited, path); found {
			fmt.Println(len(visited))
			return result, nil
		}

	}
	return nil, nil
}

func DlsFirst(currUrl string, endPage string, depth int, visited map[string]bool, path []string) (bool, []string) {
	if currUrl == endPage && depth == 0 {
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
			if found, result := DlsFirst(value, endPage, depth-1, visited, path); found {
				return true, result
			}
		}
	}

	visited[currUrl] = false // Unmark the current node
	return false, nil
}

// func IdsFirstGoRoutine(startPage string, endPage string, maxDepth int) ([]string, error) {
// 	startPage = strings.TrimSpace(startPage)
// 	endPage = strings.TrimSpace(endPage)

// 	if !scraper.IsWikiPageUrlExists(&startPage) {
// 		return nil, fmt.Errorf("start page doesn't exists")
// 	}
// 	if !scraper.IsWikiPageUrlExists(&endPage) {
// 		return nil, fmt.Errorf("end page doesn't exists")
// 	}
// 	var found int32 = 0
// 	pathsChannel := make(chan []string, 10) // assume max path is 10
// 	for depth := 0; depth < maxDepth; depth++ {
// 		var wg sync.WaitGroup
// 		visited := make(map[string]bool)
// 		path := []string{}
// 		limiter := make(chan struct{}, 5)
// 		wg.Add(1)
// 		limiter <- struct{}{}
// 		go DlsFirstGoRoutine(startPage, endPage, depth, visited, path, &found, &wg, pathsChannel, limiter)
// 		fmt.Println("Bawah dls go routine")
// 		wg.Wait()
// 		// fmt.Println("luar wait")
// 		// if found, result := DlsFirstGoRoutine(startPage, endPage, depth, visited, path, &found); found {
// 		// 	fmt.Println(len(visited))
// 		// 	return result, nil
// 		// }
// 		// finalPath, ok := <-pathsChannel
// 		// if ok {
// 		// 	return finalPath, nil
// 		// }

// 		if atomic.LoadInt32(&found) == 1 {
// 			finalPath := <-pathsChannel
// 			close(pathsChannel)
// 			return finalPath, nil
// 		}
// 		fmt.Println("luar final")

// 	}
// 	return nil, nil
// }

func IdsFirstGoRoutine(startPage string, endPage string, maxDepth int) ([]string, error) {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)

	// Check if the start and end pages exist in your scrape function
	if !scraper.IsWikiPageUrlExists(&startPage) {
		return nil, fmt.Errorf("start page doesn't exist")
	}
	if !scraper.IsWikiPageUrlExists(&endPage) {
		return nil, fmt.Errorf("end page doesn't exist")
	}

	var found int32 = 0
	pathsChannel := make(chan []string, 10) // Channel to collect valid paths
	limiter := make(chan struct{}, 50)      // Concurrency limiter

	for depth := 0; depth <= maxDepth; depth++ {
		var wg sync.WaitGroup
		visited := make(map[string]bool)
		var visitedMutex sync.Mutex
		wg.Add(1)

		go DlsFirstGoRoutine(startPage, endPage, depth, visited, []string{}, &found, &wg, pathsChannel, limiter, &visitedMutex)

		done := make(chan bool)
		go func() {
			wg.Wait()
			done <- true
		}()

		go func() {
			ticker := time.NewTicker(100 * time.Millisecond) // Check every 100ms
			defer ticker.Stop()

			for {
				select {
				case <-done:
					// All goroutines have finished
					return
				case <-ticker.C:
					if atomic.LoadInt32(&found) == 1 {
						// Found is set, you can take additional actions if necessary
						println("Target was found!")
					}
				}
			}
		}()

		// wg.Wait() // Wait for all goroutines at this depth to complete
		fmt.Println(depth)

		if atomic.LoadInt32(&found) == 1 {
			close(pathsChannel) // Ensure no more writes to the channel
			if path, ok := <-pathsChannel; ok {
				return path, nil // Return the found path
			}
		}
	}

	return nil, fmt.Errorf("no path found within depth %d", maxDepth)
}

func DlsFirstGoRoutine(currUrl string, endPage string, depth int, visited map[string]bool, path []string, found *int32, wg *sync.WaitGroup, pathsChannel chan []string, limiter chan struct{}, visitedMutex *sync.Mutex) {
	defer wg.Done()
	limiter <- struct{}{}
	// select {
	// case limiter <- struct{}{}:
	// 	// Successfully took a slot in limiter, proceed.
	// case <-time.After(100 * time.Millisecond): // Or some other sensible timeout
	// 	// Failed to take a slot in time, exit to avoid deadlock or slow processing.
	// 	return
	// }
	defer func() { <-limiter }() // Release the limiter slot when done
	// fmt.Println(currUrl)
	if currUrl == endPage {
		// fmt.Println("HERE")
		if atomic.LoadInt32(found) == 0 {
			fmt.Println("DOWN HERE")
			atomic.StoreInt32(found, 1)
			path = append(path, currUrl)
			pathsChannel <- append([]string(nil), path...)
			return
		}
		// return true, append(path, currUrl)
		return
	}
	// fmt.Println("Atas load ints")
	if depth <= 0 || atomic.LoadInt32(found) == 1 {
		// fmt.Println("masuk load ints")
		return
	}
	// fmt.Println("Atas getscrape links")
	var allUrl = scraper.GetScrapeLinksConcurrent(currUrl)
	// fmt.Println("Bawah getscrape links")
	if allUrl == nil {
		return
	}

	visitedMutex.Lock()
	visited[currUrl] = true
	visitedMutex.Unlock()
	path = append(path, currUrl)
	// limiter := make(chan struct{}, 100)
	for _, value := range allUrl {
		// limiter <- struct{}{}
		if atomic.LoadInt32(found) == 1 {
			return
		}
		visitedMutex.Lock()
		if !visited[value] {
			// fmt.Println("Atas masukin limiter")
			// limiter <- struct{}{}
			// fmt.Println("Bawah masukin limiter")
			// if found, result := DlsFirstGoRoutine(value, endPage, depth-1, visited, path, found); found {
			// 	return true, result
			// }
			wg.Add(1)
			go DlsFirstGoRoutine(value, endPage, depth-1, visited, path, found, wg, pathsChannel, limiter, visitedMutex)
		}
		visitedMutex.Unlock()
	}
	visitedMutex.Lock()
	visited[currUrl] = false // Unmark the current node
	visitedMutex.Unlock()
	// return
}

func IdsFirstGoRoutineAllPaths(startPage string, endPage string, maxDepth int) ([][]string, error) {
	startPage = strings.TrimSpace(startPage)
	endPage = strings.TrimSpace(endPage)
	var allPaths [][]string

	// Check if the start and end pages exist in your scrape function
	if !scraper.IsWikiPageUrlExists(&startPage) {
		return nil, fmt.Errorf("start page doesn't exist")
	}
	if !scraper.IsWikiPageUrlExists(&endPage) {
		return nil, fmt.Errorf("end page doesn't exist")
	}

	var found int32 = 0
	pathsChannel := make(chan []string, 10) // Channel to collect valid paths
	limiter := make(chan struct{}, 50)      // Concurrency limiter

	for depth := 0; depth <= maxDepth; depth++ {
		var wg sync.WaitGroup
		var allPathMutex sync.Mutex
		visited := make(map[string]bool)
		var visitedMutex sync.Mutex
		wg.Add(1)

		go DlsFirstGoRoutineAllPaths(startPage, endPage, depth, visited, []string{}, &found, &wg, pathsChannel, limiter, &visitedMutex, &allPaths, &allPathMutex)

		wg.Wait()

		// wg.Wait() // Wait for all goroutines at this depth to complete
		fmt.Println(depth)

		if atomic.LoadInt32(&found) == 1 {
			close(pathsChannel) // Ensure no more writes to the channel
			return allPaths, nil
		}
	}

	return nil, fmt.Errorf("no path found within depth %d", maxDepth)
}

func DlsFirstGoRoutineAllPaths(currUrl string, endPage string, depth int, visited map[string]bool, path []string, found *int32, wg *sync.WaitGroup, pathsChannel chan []string, limiter chan struct{}, visitedMutex *sync.Mutex, allPaths *[][]string, allPathMutex *sync.Mutex) {
	defer wg.Done()
	limiter <- struct{}{}
	// select {
	// case limiter <- struct{}{}:
	// 	// Successfully took a slot in limiter, proceed.
	// case <-time.After(100 * time.Millisecond): // Or some other sensible timeout
	// 	// Failed to take a slot in time, exit to avoid deadlock or slow processing.
	// 	return
	// }
	defer func() { <-limiter }() // Release the limiter slot when done
	// fmt.Println(currUrl)
	if currUrl == endPage {
		// fmt.Println("HERE")
		if atomic.LoadInt32(found) == 0 {
			fmt.Println("DOWN HERE")
			atomic.StoreInt32(found, 1)
		}
		path = append(path, currUrl)
		newPath := make([]string, len(path))
		copy(newPath, path)
		allPathMutex.Lock()
		*allPaths = append(*allPaths, newPath)
		allPathMutex.Unlock()
		fmt.Println(*allPaths)
		return
	}
	// fmt.Println("Atas load ints")
	if depth <= 0 || atomic.LoadInt32(found) == 1 {
		// fmt.Println("masuk load ints")
		return
	}
	// fmt.Println("Atas getscrape links")
	var allUrl = scraper.GetScrapeLinksConcurrent(currUrl)
	// fmt.Println("Bawah getscrape links")
	if allUrl == nil {
		return
	}

	visitedMutex.Lock()
	visited[currUrl] = true
	visitedMutex.Unlock()
	path = append(path, currUrl)
	// limiter := make(chan struct{}, 100)
	for _, value := range allUrl {
		// limiter <- struct{}{}
		visitedMutex.Lock()
		if !visited[value] {
			// fmt.Println("Atas masukin limiter")
			// limiter <- struct{}{}
			// fmt.Println("Bawah masukin limiter")
			// if found, result := DlsFirstGoRoutine(value, endPage, depth-1, visited, path, found); found {
			// 	return true, result
			// }
			wg.Add(1)
			go DlsFirstGoRoutineAllPaths(value, endPage, depth-1, visited, path, found, wg, pathsChannel, limiter, visitedMutex, allPaths, allPathMutex)
		}
		visitedMutex.Unlock()
	}
	visitedMutex.Lock()
	visited[currUrl] = false // Unmark the current node
	visitedMutex.Unlock()
	// return
}

// var mutex sync.Mutex
// var foundGlobal bool // global indicator if path is found

// func IdsFirstGoRoutine(startPage string, endPage string, maxDepth int) ([]string, error) {
// 	startPage = strings.TrimSpace(startPage)
// 	endPage = strings.TrimSpace(endPage)

// 	if !scraper.IsWikiPageUrlExists(&startPage) {
// 		return nil, fmt.Errorf("start page doesn't exists")
// 	}
// 	if !scraper.IsWikiPageUrlExists(&endPage) {
// 		return nil, fmt.Errorf("end page doesn't exists")
// 	}

// 	var result []string
// 	var wg sync.WaitGroup

// 	for depth := 0; depth < maxDepth && !foundGlobal; depth++ {
// 		visited := make(map[string]bool)
// 		path := []string{}
// 		wg.Add(1)

// 		go func(d int) {
// 			defer wg.Done()
// 			if found, res := DfsFirstGoRoutine(startPage, endPage, d, visited, path); found {
// 				mutex.Lock()
// 				if !foundGlobal { // Ensure no other goroutine has set the path
// 					result = res
// 					foundGlobal = true
// 				}
// 				mutex.Unlock()
// 			}
// 		}(depth)
// 		wg.Wait() // Wait for the goroutines of this depth level
// 		if foundGlobal {
// 			break
// 		}
// 	}

// 	return result, nil
// }

// func DfsFirstGoRoutine(currUrl string, endPage string, depth int, visited map[string]bool, path []string) (bool, []string) {
// 	if foundGlobal {
// 		return false, nil // Stop processing if the path is already found
// 	}

// 	if currUrl == endPage {
// 		return true, append(path, currUrl)
// 	}
// 	if depth <= 0 {
// 		return false, nil
// 	}

// 	mutex.Lock() // Lock before accessing global resources
// 	if visited[currUrl] {
// 		mutex.Unlock()
// 		return false, nil
// 	}
// 	visited[currUrl] = true
// 	mutex.Unlock()

// 	path = append(path, currUrl)
// 	var result []string
// 	var found bool

// 	allUrl := scraper.GetScrapeLinks(currUrl)
// 	if allUrl == nil {
// 		return false, nil
// 	}

// 	for _, value := range allUrl {
// 		if !visited[value] {
// 			if f, res := DfsFirst(value, endPage, depth-1, visited, path); f {
// 				result = res
// 				found = f
// 				break
// 			}
// 		}
// 	}

// 	mutex.Lock()
// 	visited[currUrl] = false // Unmark the current node after finishing
// 	mutex.Unlock()

// 	return found, result
// }
