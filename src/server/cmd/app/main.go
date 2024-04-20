package main

import (
	"bufio"
	"fmt"
	"os"
	algorithm "server/pkg/algorithm"
	"server/pkg/scraper"
	"time"
)

func main() {
	scraper.LoadCache()
	// defer scraper.WriteCsv("data.txt")
	defer scraper.WriteJSON("links.json")
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

	start := time.Now()

	hasil := algorithm.IdsFirstGoRoutine(urlStart, urlEnd, 5)

	end := time.Now()

	totalTime := end.Sub(start)
	fmt.Println("time : ", totalTime)

	if hasil != nil {
		fmt.Printf("Depth : %d\n", len(hasil)-1)
		fmt.Println(hasil)

	} else {
		fmt.Println("Something went wrong")
	}

}
