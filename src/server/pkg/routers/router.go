package routers

import (
	"fmt"
	"net/http"
	"server/pkg/algorithm"
	"server/pkg/scraper"
	"strconv"

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

type Properties struct {
	ID       string `json:"id"`
	DEGREES  string `json:"degrees"`
	ARTICLES string `json:"articles"`
	PATH     string `json:"path"`
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

var PropGlobal Properties

func GetResult(c *gin.Context) {
	fmt.Println("Url Start: ", searchData.URLStart)
	fmt.Println("Url Target: ", searchData.URLTarget)
	fmt.Println("Algorithm Before : ", searchData.Algorithm)
	fmt.Println("searchData.Path Before : ", searchData.Path)
	var hasil []string
	var visited map[string]bool
	var hasillMultPath [][]string
	var err error
	var counter int
	scraper.LoadCache("links.json")
	if searchData.Algorithm == "1" && searchData.Path == "1" {
		fmt.Println("Picked BFS One Solution")
		hasil, visited, err = algorithm.BfsGoRoutine(searchData.URLStart, searchData.URLTarget)

	} else if searchData.Algorithm == "2" && searchData.Path == "1" {
		fmt.Println("Picked IDS One Solution")
		hasil, counter, err = algorithm.IdsFirstPath(searchData.URLStart, searchData.URLTarget, 9)

	} else if searchData.Algorithm == "1" && searchData.Path == "2" {
		fmt.Println("Picked BFS ALL Solution")
		hasillMultPath, visited, err = algorithm.BfsAllPathGoRoutine(searchData.URLStart, searchData.URLTarget)

	} else if searchData.Algorithm == "2" && searchData.Path == "2" {
		fmt.Println("Picked IDS ALL Solution")
		hasillMultPath, counter, err = algorithm.IdsAllPath(searchData.URLStart, searchData.URLTarget, 10)

	}
	fmt.Println("Algorithm after: ", searchData.Algorithm)
	fmt.Println("searchData.Path after : ", searchData.Path)

	if err != nil {
		return
	}
	if searchData.Path == "1" {
		fmt.Println(hasil)

		data := make([]Result, 1) // ini masih satu path saja
		data[0] = Result{
			ID:    "1",
			Title: scraper.PathToTitle(hasil),
			URL:   hasil,
		}
		PropGlobal.DEGREES = strconv.Itoa(len(hasil) - 1)
		PropGlobal.PATH = "1"
		c.IndentedJSON(http.StatusCreated, data)

	} else if searchData.Path == "2" {

		results := make([]Result, len(hasillMultPath)) // berapa banyak path
		for id, hasil := range hasillMultPath {
			results[id] = Result{
				ID:    strconv.Itoa(id + 1),
				Title: scraper.PathToTitle(hasillMultPath[id]),
				URL:   hasil,
			}
		}
		c.IndentedJSON(http.StatusCreated, results)
		PropGlobal.DEGREES = strconv.Itoa(len(hasillMultPath[0]) - 1)
		PropGlobal.PATH = strconv.Itoa(len(hasillMultPath))
	}
	PropGlobal.ID = "1"
	if searchData.Algorithm == "1" {
		PropGlobal.ARTICLES = strconv.Itoa(len(visited))
	} else if searchData.Algorithm == "2" {
		PropGlobal.ARTICLES = strconv.Itoa(counter)
	}

	fmt.Println("Number of paths : ", PropGlobal.PATH)
	fmt.Println("DEPTH : ", PropGlobal.DEGREES)
	fmt.Println("AMOUNT OF ARTICLES: ", PropGlobal.ARTICLES)

}

func GetProperties(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, PropGlobal)
}

func OtherFunction() {
	fmt.Println("Search Start in OtherFunction:", searchData.URLStart)
}
