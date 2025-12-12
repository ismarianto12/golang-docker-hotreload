package main

import (
	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	Port string
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"data":    "rian",
			"message": "perubahan",
		})
	})

	r.Group("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API v1",
		})
	})
	r.Run(":6060") // listen and serve on

}
