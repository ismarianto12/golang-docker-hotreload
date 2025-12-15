package routes

import (
	"rianRestapp/usecases"

	"github.com/gin-gonic/gin"
)

func IntialRoute(port string) {
	r := gin.Default()
	prod := usecases.NewProductUsecase()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			product := v1.Group("/product")
			{
				product.GET("/index", prod.GetProductInfo)
				product.GET("/list", prod.GetIndexData)
				product.POST("/create", prod.CreateProd)
				product.POST("/insert", prod.UpdateData)
			}

		}
	}

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
	r.Run(port)

}
