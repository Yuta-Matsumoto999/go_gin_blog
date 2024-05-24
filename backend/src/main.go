package main

import (
	"github.com/gin-gonic/gin"

	"src/lib/database"
)

func main() {
	r := gin.Default()

	db := database.DBConnection()

	defer db.Close()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!!",
		})
	})

	r.GET("/other", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Wooooo!!!",
		})
	})
	r.Run(":8080")
}
