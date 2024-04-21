package main

import (
	router1 "server/pkg/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	defer scraper.WriteJSON("links.json") // disable cache for real world testing
// 	scraper.LoadCache()
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Start Page Title : ")
// 	urlStart, err := reader.ReadString('\n')
// 	if err != nil {
// 		return
// 	}
// 	urlStart = scraper.TitleToWikiUrl(urlStart)
// 	fmt.Print("End Page Title : ")
// 	urlEnd, err := reader.ReadString('\n')
// 	if err != nil {
// 		return
// 	}
// 	urlEnd = scraper.TitleToWikiUrl(urlEnd)
// 	fmt.Println(urlStart)
// 	fmt.Println(urlEnd)
// 	fmt.Print("Algorithm : \n")
// 	fmt.Print("1.BFS \n")
// 	fmt.Print("2.IDS \n")
// 	fmt.Print("3.BFS Go Routine \n")
// 	fmt.Print("4.IDS Go Routine \n")
// 	fmt.Print("Input : ")
// 	var algo_input int
// 	_, err = fmt.Scan(&algo_input)
// 	if err != nil {
// 		return
// 	}
// 	var hasil []string
// 	start := time.Now()
// 	if algo_input == 1 {
// 		hasil, err = algorithm.Bfs(urlStart, urlEnd)

// 	} else if algo_input == 2 {
// 		hasil, err = algorithm.IdsFirst(urlStart, urlEnd, 10)

// 	} else if algo_input == 3 {
// 		// hasil, err = algorithm.BfsGoRoutine(urlStart, urlEnd)
// 	} else if algo_input == 4 {
// 		hasil, err = algorithm.IdsFirstGoRoutine(urlStart, urlEnd, 10)
// 	} else {
// 		return
// 	}
// 	end := time.Now()

// 	if err != nil {
// 		fmt.Println(err)
// 	} else if hasil == nil {
// 		fmt.Println("No possible path")
// 	} else {
// 		fmt.Println("Time : ", end.Sub(start))
// 		fmt.Printf("Depth : %d\n", len(hasil))
// 		fmt.Println(hasil)
// 		fmt.Println(scraper.PathToTitle(hasil))
// 	}

// }

func main() {
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Routes
	router.POST("/api/search", router1.GetSearch)
	router.GET("/api/result", router1.GetResult)
	// router.POST("api/properties", GetProperties)

	// Run the server
	router.Run("localhost:8080")
}
