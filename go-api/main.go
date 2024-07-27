package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("Added router")

	// Define a route for the GET method
	r.GET("/hello", func(c *gin.Context) {
		// response := map[string]string{
		// 	"message": "Hello, World!",
		// }
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Define a route for the POST method
	r.POST("/greet", func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message": "Hello, " + json.Name,
		})
	})

	// Run the server on port 8080
	r.Run(":8081")
}
