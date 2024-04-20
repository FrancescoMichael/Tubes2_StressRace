package main

import (
	"bufio"
	"fmt"
	"os"
	algorithm "server/pkg/algorithm"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Url start page : ")
	urlStart, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	}
	fmt.Print("Url end page : ")
	urlEnd, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
	}
	// hasil, depth := algorithm.Ids(urlStart, urlEnd, 100)

	// fmt.Print("Depth : ", depth, "\n")
	// for _, var2 := range hasil {
	// 	fmt.Println(var2)
	// }

	hasil := algorithm.IdsFirst(urlStart, urlEnd, 5)

	if hasil != nil {
		fmt.Printf("Depth : %d \n", len(hasil)-1)
		fmt.Println(hasil)
	} else {
		fmt.Println("No possible solution")
	}

	// router := gin.Default()

	// router.Use(cors.Default())

	// router.GET("/albums", routers.GetAlbums)
	// router.POST("/albums", routers.PostAlbums)

	// router.Run("localhost:8080")

}
