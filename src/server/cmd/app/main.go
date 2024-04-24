package main

import (
	routers "server/pkg/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Routes
	router.POST("/api/search", routers.GetSearch)
	router.GET("/api/result", routers.GetResult)
	// router.POST("api/properties", GetProperties)

	// Run the server
	router.Run(":8080")
}
