package test

import (
	"fmt"
	"server/pkg/algorithm"
	"server/pkg/scraper"
	"time"
)

func Test() {
	// defer scraper.WriteJSON("links.json")
	// scraper.LoadCache()
	// reader := bufio.NewReader(os.Stdin)
	fmt.Print("Start Page Title : ")
	// urlStart := "https://en.wikipedia.org/wiki/Joko_Widodo"
	urlStart := "https://en.wikipedia.org/wiki/Neuroscience"
	// urlStart := "https://en.wikipedia.org/wiki/Russia"
	// urlStart := "https://en.wikipedia.org/wiki/Neuroscience"
	// if err != nil {
	// 	return
	// }
	// urlStart = scraper.TitleToWikiUrl(urlStart)
	fmt.Print("End Page Title : ")
	// urlEnd, err := reader.ReadString('\n')
	// urlEnd := "https://en.wikipedia.org/wiki/Atheism"
	urlEnd := "https://en.wikipedia.org/wiki/Springtail"
	// urlEnd := "https://en.wikipedia.org/wiki/Joko_Widodo"
	// urlEnd := "https://en.wikipedia.org/wiki/Brain%E2%80%93computer_interface"
	// if err != nil {
	// 	return
	// }
	// urlEnd = scraper.TitleToWikiUrl(urlEnd)
	fmt.Println(urlStart)
	fmt.Println(urlEnd)
	fmt.Print("Algorithm : \n")
	fmt.Print("1.BFS \n")
	fmt.Print("2.IDS \n")
	fmt.Print("3.BFS Go Routine \n")
	fmt.Print("4.IDS Go Routine \n")
	fmt.Print("5.BFS mult path \n")
	fmt.Print("6.IDS mult path \n")
	fmt.Print("Input : ")
	var algo_input int = 4
	var visited map[string]bool

	var hasil []string
	var allPath [][]string
	var err error
	start := time.Now()
	if algo_input == 1 {
		// hasil, err = algorithm.Bfs(urlStart, urlEnd)
		hasil, visited, err = algorithm.BfsGoRoutine(urlStart, urlEnd)
	} else if algo_input == 2 {
		// hasil, visited, err = algorithm.IdsFirstGoRoutine(urlStart, urlEnd, 10)
		// hasil, visited, err = algorithm.IdsFirstGoRoutineExp(urlStart, urlEnd, 10)
		// hasil, visited, err = algorithm.IdsFirstGoRoutine(urlStart, urlEnd, 10)

	} else if algo_input == 3 {
		hasil, visited, err = algorithm.BfsGoRoutine(urlStart, urlEnd)
	} else if algo_input == 4 {
		hasil, visited, err = algorithm.IdsFirstPath(urlStart, urlEnd, 10)
	} else if algo_input == 5 {

		allPath, visited, err = algorithm.BfsAllPathGoRoutine(urlStart, urlEnd)
	} else if algo_input == 6 {
		allPath, visited, _ = algorithm.IdsAllPath(urlStart, urlEnd, 10)
	}
	end := time.Now()
	if err != nil {
		fmt.Println(err)
	} else if hasil == nil && allPath == nil {
		fmt.Println("No possible path")
	} else {
		fmt.Println("Time : ", end.Sub(start))
		if algo_input == 1 || algo_input == 2 || algo_input == 3 || algo_input == 4 {
			fmt.Printf("Depth : %d\n", len(hasil))
			fmt.Printf("Articles Checked : %d\n", len(visited))
			fmt.Println(hasil)
			fmt.Println(scraper.PathToTitle(hasil))
		} else {
			for _, link := range allPath {
				fmt.Println(link)
			}
			fmt.Printf("Articles Checked : %d\n", len(visited))
			fmt.Printf("Depth : %d\n", len(allPath[0]))
			fmt.Printf("Amount of path : %d\n", len(allPath))
		}
	}

}
