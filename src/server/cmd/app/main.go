package main

import (
	"server/pkg/routers"

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
	router.GET("/api/properties", routers.GetProperties)

	// Run the server
	router.Run(":8080")
}

// package main

// import (
// 	"log"
// 	"os"
// 	"os/signal"
// 	"server/pkg/test"
// 	"syscall"
// )

// func main() {
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	go func() {
// 		sig := <-sigs
// 		log.Printf("Received signal: %s", sig)
// 		// Perform cleanup tasks, close resources, etc.
// 		os.Exit(0) // Exit successfully after handling
// 	}()
// 	test.Test()
// }
