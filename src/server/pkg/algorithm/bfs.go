package algorithm

import (
	scraper "server/scraper"
	"fmt"
)

type Graph map[string][]string

func (g *Graph) AddVertex(vertex string) {
	(*g)[vertex] = []string{}
}

func (g *Graph) AddEdge(from, to string) {
	(*g)[from] = append((*g)[from], to)
	// (*g)[to] = append((*g)[to], from)
}

func bfs(startPage string, endPage string) ([][]string) {
	var ans [][]string
	ans = append(ans, shortestPath[0])
	if len(shortestPath) > 1 {
		for i := 1; i <len(shortestPath); i++ {
			if(len(shortestPath[i]) == len(shortestPath[0])) {
				ans = append(ans, shortestPath[i])
			}
		}
	}

	return ans
}

func ShortestPaths(graph Graph, start, end string) [][]string {
	// Initialize variables
	paths := make([][]string, 0)
	queue := [][]string{{start}}
	visited := make(map[string]bool)

	// Breadth-first search
	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		node := currentPath[len(currentPath)-1]

		if node == end {
			// If the current node is the end node, append the path to the result
			paths = append(paths, currentPath)
			continue
		}

		// Mark the current node as visited
		visited[node] = true

		// Explore neighbors
		for _, neighbor := range graph[node] {
			// Check if the neighbor has not been visited
			if !visited[neighbor] {
				// Add the neighbor to the current path and enqueue it
				newPath := append(append([]string{}, currentPath...), neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	// Return the found paths
	return paths 
}

func main() {

	var n, m int
	fmt.Print("Number of node : ")
	fmt.Scanln(&n)
	fmt.Print("Number of edge : ")
	fmt.Scanln(&m)

	graph := make(Graph)

	for i := 0; i < n; i++ {
		var input string
		fmt.Scanln(&input)
		graph.AddVertex(input)
	}

	for i:= 0; i < m; i++ {
		var start, end string
		fmt.Scanln(&start, &end)
		graph.AddEdge(start, end)
	}

	var startNode, endNode string
	fmt.Scanln(&startNode, &endNode)

	shortestPath := ShortestPaths(graph, startNode, endNode)
	// fmt.Println("Shortest Path:", shortestPath)
	var ans [][]string
	ans = append(ans, shortestPath[0])
	if len(shortestPath) > 1 {
		for i := 1; i <len(shortestPath); i++ {
			if(len(shortestPath[i]) == len(shortestPath[0])) {
				ans = append(ans, shortestPath[i])
			}
		}
	}

	fmt.Println("Shortest Path:", ans)
	

	// allPaths := AllPaths(graph, "A", "G")
	// fmt.Println("All Paths:", allPaths)
}

func AllPaths(graph Graph, start, end string) [][]string {
	var paths [][]string
	var currentPath []string
	visited := make(map[string]bool)

	findPaths(graph, start, end, visited, currentPath, &paths)
	return paths
}

func findPaths(graph Graph, node, end string, visited map[string]bool, currentPath []string, paths *[][]string) {
	currentPath = append(currentPath, node)
	visited[node] = true

	if node == end {
		*paths = append(*paths, append([]string{}, currentPath...))
	} else {
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				findPaths(graph, neighbor, end, visited, currentPath, paths)
			}
		}
	}

	delete(visited, node)
}
