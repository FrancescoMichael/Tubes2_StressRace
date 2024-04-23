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
	Path      string `json:"path"`
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
	fmt.Println(searchData.URLStart)
	var hasil []string
	// var hasillMultPath [][]string
	var err error
	scraper.LoadCache()
	if searchData.Algorithm == "1" && searchData.Path == "1" {
		hasil, err = algorithm.Bfs(searchData.URLStart, searchData.URLTarget)

	} else if searchData.Algorithm == "2" && searchData.Path == "1" {
		hasil, err = algorithm.IdsFirst(searchData.URLStart, searchData.URLTarget, 9)
	}
	// } else if searchData.Algorithm == "1" && searchData.Path == "2" {
	// 	hasillMultPath, err = algorithm.IdsFirstGoRoutineAllPaths(searchData.URLStart, searchData.URLTarget, 9)

	// } else if searchData.Algorithm == "2" && searchData.Path == "2" {

	// }
	// handle ketika ada error belom ada
	if err != nil {
		return
	}
	data := make([]Result, 1) // ini masih satu path saja

	fmt.Println("Ini hasil : ", hasil)

	// data[0] = Result{
	// 	ID:    "1",
	// 	Title: scraper.PathToTitle(hasil),
	// 	URL:   hasil,
	// }

	c.IndentedJSON(http.StatusCreated, data)

	fmt.Println("halo")
}

func OtherFunction() {
	fmt.Println("Search Start in OtherFunction:", searchData.URLStart)
}
