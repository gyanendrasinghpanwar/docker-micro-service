package main

import (
	"fmt"

	"my-go-project/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("Added router")

	registry.RegisterRoutes(r)

	// Run the server on port 8081
	r.Run(":8081")
}
