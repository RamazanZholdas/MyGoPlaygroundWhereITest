package main

import "github.com/gin-gonic/gin"

const (
	port = "8080"
)

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})

	router.Run(":" + port)
}
