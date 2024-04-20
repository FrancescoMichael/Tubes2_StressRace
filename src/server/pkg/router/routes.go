package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetSearch(c *gin.Context) {
    var searchData struct {
        SearchStart  string `json:"searchStart"`
        SearchTarget string `json:"searchTarget"`
        Algorithm    string `json:"algorithm"`
    }

    if err := c.BindJSON(&searchData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Search Start:", searchData.SearchStart)
    fmt.Println("Search Target:", searchData.SearchTarget)
    fmt.Println("Algorithm:", searchData.Algorithm)

    c.JSON(http.StatusOK, searchData)
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
	router.POST("/search", GetSearch)

	// Run the server
	router.Run("localhost:8080")
}
