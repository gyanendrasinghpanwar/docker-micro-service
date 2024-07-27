package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"target" binding:"required"`
}

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

	r.POST("/binary-search", func(c *gin.Context) {
		var requestBody RequestBody

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		numbers := requestBody.Numbers
		index := binarySearch(numbers, requestBody.Target)

		c.JSON(200, gin.H{
			"index": index,
		})
	})

	// Run the server on port 8081
	r.Run(":8081")
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Target not found
}
