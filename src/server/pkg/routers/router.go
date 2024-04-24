package routers

import (
	"fmt"
	"net/http"

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
	// defer scraper.WriteJSON("links.json")
	// fmt.Println(searchData.URLStart)
	// var hasil []string
	// var visited [string]bool
	// var hasillMultPath [][]string
	// var err error
	// scraper.LoadCache()
	// if searchData.Algorithm == "1" && searchData.Path == "1" {
	// 	hasil, visited, err = algorithm.BfsGoRoutine(searchData.URLStart, searchData.URLTarget)

	// } else if searchData.Algorithm == "2" && searchData.Path == "1" {
	// 	hasil, visited, err = algorithm.IdsFirst(searchData.URLStart, searchData.URLTarget, 9)

	// } else if searchData.Algorithm == "1" && searchData.Path == "2" {
	// 	hasillMultPath, visited, err = algorithm.BfsMultPath(searchData.URLStart, searchData.URLTarget)

	// } else if searchData.Algorithm == "2" && searchData.Path == "2" {
	// 	hasillMultPath, visited, err = algorithm.IdsFirstGoRoutineAllPaths(searchData.URLStart, searchData.URLTarget, 10)

	// }
	// // handle ketika ada error belom ada
	// if err != nil {
	// 	return
	// }
	// if searchData.Path == "1" {
	// 	data := make([]Result, 1) // ini masih satu path saja
	// 	data[0] = Result{
	// 		ID:    "1",
	// 		Title: scraper.PathToTitle(hasil),
	// 		URL:   hasil,
	// 	}
	// 	// 	c.IndentedJSON(http.StatusCreated, data)

	// } else if searchData.Path == "2" {
	// 	data := make()
	// }

	// // 	fmt.Println("Ini hasil : ", hasil)

	// // fmt.Println("halo")
}

func OtherFunction() {
	fmt.Println("Search Start in OtherFunction:", searchData.URLStart)
}
