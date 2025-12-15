package routes

import (
	"rianRestapp/usecases"

	"github.com/gin-gonic/gin"
)

func IntialRoute(port string) {
	r := gin.Default()
	prod := usecases.NewProductUsecase()
	typeProd := usecases.NewTypeBarangUses()

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
			category := v1.Group("/category")
			{
				category.GET("/list", typeProd.Alldata)
				category.POST("/insert", typeProd.InserData)

				category.GET("/index", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"data": "status", "message": "successfully load data",
					})
				})
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
