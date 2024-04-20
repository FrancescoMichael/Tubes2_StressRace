package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	algorithm "server/pkg/algorithm"
	"server/pkg/scraper"
)

func main() {
	scraper.LoadCache()
	defer scraper.WriteCsv("data.txt")
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Url start page : ")
	urlStart, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	}
	fmt.Print("Url end page : ")
	urlEnd, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	}

	var choice string
	fmt.Print("IDS/BFS(1/2) : ")
	fmt.Scanln(&choice)

	currentTime := time.Now()

	if choice == "1" {
		hasil, depth := algorithm.Ids(urlStart, urlEnd, 100)

		fmt.Print("Depth : ", depth, "\n")
		for _, var2 := range hasil {
			fmt.Println(var2)
		}
	} else {
		hasil := algorithm.bfs(urlStart, urlEnd)
		fmt.Println("BFS")
		for _, var2 := range hasil { 
			fmt.Println(var2)
		}
	}

	endTime := time.Now()
	diff := endTime.Sub(currentTime)
	fmt.Println("Duration : ", diff)


	

}
