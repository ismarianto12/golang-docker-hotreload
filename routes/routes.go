package routes

import (
	"log"
	_ "rianRestapp/docs"
	"rianRestapp/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func IntialRoute(port string) {
	r := gin.Default()
	uc := NewUsecases()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			product := v1.Group("/product")
			{
				product.GET("/list", uc.Product.GetIndexData)
				product.POST("/create", uc.Product.CreateProd)
				product.POST("/insert", uc.Product.UpdateData)
			}

			suplier := v1.Group("/suplier")
			{
				suplier.GET("/index", uc.Suplier.IndexData)
				suplier.POST("/create", uc.Suplier.Create)
				suplier.POST("/update/:id", uc.Suplier.UpdateSuplier)
				suplier.POST("/delete/:id", uc.Suplier.Delete)
				suplier.GET("/show/:id", uc.Suplier.ShowById)
				suplier.POST("/uploadfile", uc.Suplier.UpdateDataImage)
			}
			userroute := v1.Group("/users")
			{
				userroute.GET("/index", uc.UserUseCase.IndexData)
			}

			stockmovement := v1.Group("/stockmovement")
			{
				stockmovement.GET("/index", uc.StockMovement.IndeXalldata)
				stockmovement.POST("/insert", uc.StockMovement.CreatedData)
				stockmovement.POST("/callapi", uc.StockMovement.CallApi)
				stockmovement.POST("/testdaa", uc.StockMovement.TestPostData)
			}

			category := v1.Group("/category")
			{
				category.GET("/list", handlers.CheckTokenHeader(), uc.TypeBarang.Alldata)
				category.POST("/insert", uc.TypeBarang.InserData)
				category.POST("/show/:id", func(ctx *gin.Context) {
					id := ctx.Param("id")
					num, err := strconv.Atoi(id)
					if err != nil {
						ctx.JSON(400, gin.H{
							"id":   num,
							"data": "error invalid param id",
						})
						return
					}

					log.Printf("logging id %d", num)
					ctx.JSON(200, gin.H{
						"id":      num,
						"status":  "success",
						"message": "successfully load data",
					})
				})
			}
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "Route not found",
			"path":    c.Request.URL.Path,
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"data":    "rian",
			"message": "perubahan",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	r.Run(port)
}
