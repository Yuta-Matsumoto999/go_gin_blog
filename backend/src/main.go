package main

import (
	"src/lib/database"

	"github.com/gin-gonic/gin"
)

func init() {
	db := database.DBConnection()
	database.DBMigration(db)
}

func main() {
	r := gin.Default()

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
