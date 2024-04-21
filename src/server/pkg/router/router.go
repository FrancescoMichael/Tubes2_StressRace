package router

import (
	"fmt"
	"net/http"

	algorithm "server/pkg/algorithm"
	scraper "server/pkg/scraper"

	"github.com/gin-gonic/gin"
	// "../../cmd/app/main"
)

// Data structure
// search input
type SearchData struct {
	URLStart  string `json:"urlStart"`
	URLTarget string `json:"urlTarget"`
	Algorithm string `json:"algorithm"`
}

// result
type Result struct {
	ID    string   `json:"id"`
	Title []string `json:"title"`
	URL   []string `json:"url"`
}

var searchData SearchData

func GetSearch(c *gin.Context) {
	if err := c.BindJSON(&searchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OtherFunction()

	c.JSON(http.StatusOK, searchData)
}

func GetResult(c *gin.Context) {
	defer scraper.WriteJSON("links.json")
	// searchData.URLStart = "https://en.wikipedia.org/wiki/Adolf_Hitler"
	// searchData.URLTarget = "https://en.wikipedia.org/wiki/Nazi_Germany"
	// searchData.Algorithm = "1"
	fmt.Println(searchData.URLStart)
	var hasil []string
	var err error
	scraper.LoadCache()
	if searchData.Algorithm == "1" {
		hasil, err = algorithm.Bfs(searchData.URLStart, searchData.URLTarget)

	} else if searchData.Algorithm == "2" {
		hasil, err = algorithm.IdsFirst(searchData.URLStart, searchData.URLTarget, 9)

	}

	// handle ketika ada error belom ada
	if err != nil {
		return
	}
	data := make([]Result, 1) // ini masih satu path saja
	// data := []Result{
	// 	{
	// 		ID: "1",
	// 		Title: []string{
	// 			"Hampi",
	// 			"Hampi (town)",
	// 			"Hampi Express",
	// 			"Michael Jordan",
	// 		},
	// 		URL: []string{
	// 			"https://en.wikipedia.org/wiki/Hampi",
	// 			"https://en.wikipedia.org/wiki/Hampi_(town)",
	// 			"https://en.wikipedia.org/wiki/Hampi_Express",
	// 			"https://en.wikipedia.org/wiki/Michael_Jordan",
	// 		},
	// 	},
	// 	{
	// 		ID: "2",
	// 		Title: []string{
	// 			"Michael",
	// 			"Michael Jackson",
	// 			"Michael Jordan",
	// 			"Michael Jordan",
	// 		},
	// 		URL: []string{
	// 			"https://en.wikipedia.org/wiki/Michael",
	// 			"https://en.wikipedia.org/wiki/Michael_Jackson",
	// 			"https://en.wikipedia.org/wiki/Michael_Jordan",
	// 			"https://en.wikipedia.org/wiki/Michael_Jordan",
	// 		},
	// 	},
	// }

	fmt.Println("Ini hasil : ", hasil)

	data[0] = Result{
		ID:    "1",
		Title: scraper.PathToTitle(hasil),
		URL:   hasil,
	}

	c.IndentedJSON(http.StatusCreated, data)

	fmt.Println("halo")
}

// func GetProperties(c *gin.Context) {
// 	c.IndentedJSON(http.StatusCreated, main.properties())
// }

func OtherFunction() {
	fmt.Println("Search Start in OtherFunction:", searchData.URLStart)
}
