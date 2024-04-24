package main

import "server/pkg/test"

func main() {
	// router := gin.Default()

	// // CORS middleware
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type"},
	// 	AllowCredentials: true,
	// }))

	// // Routes
	// router.POST("/api/search", routers.GetSearch)
	// router.GET("/api/result", routers.GetResult)

	// // Run the server
	// router.Run("localhost:8080")
	test.Test()
}
