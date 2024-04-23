package test

import (
	"bufio"
	"fmt"
	"os"
	"server/pkg/algorithm"
	"server/pkg/scraper"
	"time"
)

func Test() {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= 1; i++ {
		fmt.Print("Start Page Title : ")
		urlStart, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		urlStart = scraper.TitleToWikiUrl(urlStart)
		fmt.Print("End Page Title : ")
		urlEnd, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		urlEnd = scraper.TitleToWikiUrl(urlEnd)
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
		var algo_input int
		// var placeholder int
		_, err = fmt.Scan(&algo_input)
		if err != nil {
			return
		}
		reader.ReadString('\n')
		var hasil []string
		var allPath [][]string
		start := time.Now()
		if algo_input == 1 {
			// hasil, err = algorithm.Bfs(urlStart, urlEnd)
			hasil, err = algorithm.BfsGoRoutine(urlStart, urlEnd)
		} else if algo_input == 2 {
			hasil, err = algorithm.IdsFirst(urlStart, urlEnd, 10)

		} else if algo_input == 3 {
			hasil, err = algorithm.BfsGoRoutine(urlStart, urlEnd)
		} else if algo_input == 4 {
			hasil, err = algorithm.IdsFirstGoRoutine(urlStart, urlEnd, 10)
		} else if algo_input == 5 {
			allPath, err = algorithm.BfsMultPath(urlStart, urlEnd)
		} else if algo_input == 6 {
			allPath, _ = algorithm.IdsFirstGoRoutineAllPaths(urlStart, urlEnd, 10)
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
				fmt.Println(hasil)
				fmt.Println(scraper.PathToTitle(hasil))

			} else {

				for _, link := range allPath {
					fmt.Println(link)
				}
				fmt.Printf("Depth : %d\n", len(allPath[0]))
				fmt.Printf("Amount of path : %d\n", len(allPath))
			}

		}

		scraper.LinkCache = make(map[string][]string)

	}

}
