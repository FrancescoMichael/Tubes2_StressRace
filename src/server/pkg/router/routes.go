package main

import (
    "fmt"
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
	// "../../cmd/app/main"
)

// Definisikan struktur data untuk menyimpan data pencarian
type SearchData struct {
    URLStart  string `json:"urlStart"`
    URLTarget string `json:"urlTarget"`
    Algorithm    string `json:"algorithm"`
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

// func GetResult(c *gin.Context) {
// 	c.IndentedJSON(http.StatusCreated, main.result(searchData.SearchStart, searchData.SearchTarget, searchData.Algorithm))
// }

// func GetProperties(c *gin.Context) {
// 	c.IndentedJSON(http.StatusCreated, main.properties())
// }

func OtherFunction() {
    fmt.Println("Search Start in OtherFunction:", searchData.URLStart)
}

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
    router.POST("/api/search", GetSearch)
	// router.POST("/api/result", GetResult)
	// router.POST("api/properties", GetProperties)

    // Run the server
    router.Run("localhost:8080")
}
