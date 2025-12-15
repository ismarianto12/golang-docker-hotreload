package routes

import (
	"log"
	"rianRestapp/handlers"
	"rianRestapp/usecases"
	"strconv"

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
				category.GET("/list", handlers.CheckTokenHeader(), typeProd.Alldata)
				category.POST("/insert", typeProd.InserData)
				category.POST("/show/:id", func(ctx *gin.Context) {
					id := ctx.Param("id")
					num, err := strconv.Atoi(id)
					if err != nil {
						log.Println("logging id %s", num)
						ctx.JSON(400, gin.H{
							"id":   num,
							"data": "error ivalid paramid",
						})
						return
					}

					log.Println("logging id %s", id)

					ctx.JSON(200, gin.H{
						"id":   num,
						"data": "status", "message": "successfully load data",
					})

				})

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
