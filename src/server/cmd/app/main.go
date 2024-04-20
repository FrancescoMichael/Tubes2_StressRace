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
	defer scraper.WriteJSON("data.json")
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
	// hasil, depth := algorithm.Ids(urlStart, urlEnd, 100)

	// fmt.Print("Depth : ", depth, "\n")
	// for _, var2 := range hasil {
	// 	fmt.Println(var2)
	// }
	start := time.Now()

	hasil := algorithm.IdsFirstGoRoutine(urlStart, urlEnd, 5)

	end := time.Now()

	totalTime := end.Sub(start)
	fmt.Println("time : ", totalTime)

	if hasil != nil {
		fmt.Printf("Depth : %d \n", len(hasil)-1)

		fmt.Println(hasil)

	} else {
		fmt.Println("No possible solution")
	}

}
