package main

import (
	"rianRestapp/usecases"

	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	Port string
}

func main() {
	r := gin.Default()
	prod := usecases.NewProductUsecase(nil)
	r.GET("/product/list", prod.GetProductInfo)
	r.GET("/product/index", prod.GetIndexData)
	r.POST("/product/insert", prod.CreateProd)

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
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Group("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API v1",
		})
	})
	r.Run(":6060") // listen and serve on

}
