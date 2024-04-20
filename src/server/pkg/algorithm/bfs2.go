package algorithm

import (
	"fmt"
	scraper "server/pkg/scraper"
)

type Graph map[string][]string

func (g *Graph) AddVertex(vertex string) {
	(*g)[vertex] = []string{}
}

func (g *Graph) AddEdge(from, to string) {
	(*g)[from] = append((*g)[from], to)
	// (*g)[to] = append((*g)[to], from)
}

// func bfs(startPage string, endPage string) ([][]string) {
// 	graph := make(Graph)

// 	shortestPath := BFSShortestPaths(graph, startPage, endPage)

// 	var ans [][]string
// 	ans = append(ans, shortestPath[0])
// 	if len(shortestPath) > 1 {
// 		for i := 1; i <len(shortestPath); i++ {
// 			if(len(shortestPath[i]) == len(shortestPath[0])) {
// 				ans = append(ans, shortestPath[i])
// 			}
// 		}
// 	}

// 	return ans
// }

// func BFSShortestPaths(graph Graph, start, end string) [][]string {
// all paths
// // Initialize variables
// paths := make([][]string, 0)
// queue := [][]string{{start}}
// visited := make(map[string]bool)

// // Breadth-first search
// for len(queue) > 0 {
// 	currentPath := queue[0]
// 	queue = queue[1:]
// 	node := currentPath[len(currentPath)-1]

// 	if node == end {
// 		// If the current node is the end node, append the path to the result
// 		paths = append(paths, currentPath)
// 		break
// 	}

// 	// Mark the current node as visited
// 	visited[node] = true
// 	var allUrl = scraper.GetScrapeLinks(node)
// 	fmt.Println(node, " : ", len(allUrl))
// 	for i := 0; i < len(allUrl); i++ {
// 		if _, exists := graph[allUrl[i]]; !exists {
// 			graph.AddVertex(allUrl[i])
// 			graph.AddEdge(node, allUrl[i])
// 		}
// 	}

// 	// Explore neighbors
// 	for _, neighbor := range graph[node] {
// 		// Check if the neighbor has not been visited
// 		if !visited[neighbor] {
// 			// Add the neighbor to the current path and enqueue it
// 			newPath := append(append([]string{}, currentPath...), neighbor)
// 			queue = append(queue, newPath)
// 		}
// 	}
// }

// // Return the found paths
// return paths
// }

func ShortestPath(graph Graph, start, end string) []string {
	queue := []string{start}
	visited := make(map[string]bool)
	parent := make(map[string]string)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node] = true
		var allUrl = scraper.GetScrapeLinks(node)
		fmt.Println(node, " : ", len(allUrl))
		for i := 0; i < len(allUrl); i++ {
			if _, exists := graph[allUrl[i]]; !exists {
				graph.AddVertex(allUrl[i])
				graph.AddEdge(node, allUrl[i])
			}
		}

		if node == end {
			path := []string{end}
			for parent[node] != start {
				path = append([]string{parent[node]}, path...)
				node = parent[node]
			}
			path = append([]string{start}, path...)
			return path
		}

		neighbors := graph[node] // Get neighbors of the current node

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				queue = append(queue, neighbor) // Add neighbor to the queue
				visited[neighbor] = true
				parent[neighbor] = node
			}
		}
	}

	return nil // no path found
}

func SingleShortestPath(startPage, endPage string) []string {
	graph := make(Graph)
	shortestPath := ShortestPath(graph, startPage, endPage)

	return shortestPath
}

// func main() {
// 	var startUrl, endUrl string
// 	fmt.Print("Start URL : ")
// 	fmt.Scanln(&startUrl)

// 	fmt.Print("End URL : ")
// 	fmt.Scanln(&endUrl)

// 	// var ans [][]string
// 	// ans = bfs(startUrl, endUrl)

// 	var ans []string
// 	ans = SingleShortestPath(startUrl, endUrl)

// 	fmt.Println(ans)
// 	// var n, m int
// 	// fmt.Print("Number of node : ")
// 	// fmt.Scanln(&n)
// 	// fmt.Print("Number of edge : ")
// 	// fmt.Scanln(&m)

// 	// graph := make(Graph)

// 	// for i := 0; i < n; i++ {
// 	// 	var input string
// 	// 	fmt.Scanln(&input)
// 	// 	graph.AddVertex(input)
// 	// }

// 	// for i:= 0; i < m; i++ {
// 	// 	var start, end string
// 	// 	fmt.Scanln(&start, &end)
// 	// 	graph.AddEdge(start, end)
// 	// }

// 	// var startNode, endNode string
// 	// fmt.Scanln(&startNode, &endNode)

// 	// shortestPath := ShortestPaths(graph, startNode, endNode)
// 	// // fmt.Println("Shortest Path:", shortestPath)
// 	// var ans [][]string
// 	// ans = append(ans, shortestPath[0])
// 	// if len(shortestPath) > 1 {
// 	// 	for i := 1; i <len(shortestPath); i++ {
// 	// 		if(len(shortestPath[i]) == len(shortestPath[0])) {
// 	// 			ans = append(ans, shortestPath[i])
// 	// 		}
// 	// 	}
// 	// }

// 	// fmt.Println("Shortest Path:", ans)
// }
