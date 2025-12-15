package usecases

import (
	"fmt"
	"log"
	"rianRestapp/entities"
	"rianRestapp/repositories"

	"github.com/gin-gonic/gin"
)

type ProductUsecase struct {
	barangRepo *repositories.BarangRepo
}

func NewProductUsecase() *ProductUsecase {
	barangRepo := repositories.NewBarangRepo()
	return &ProductUsecase{barangRepo: barangRepo}
}

func (u *ProductUsecase) GetIndexData(c *gin.Context) {
	if err, data := u.barangRepo.GetAllData(); err != nil {
		c.JSON(400, gin.H{
			"data":    err,
			"message": "success",
		})
		return
	} else {
		c.JSON(400, gin.H{
			"data":    data,
			"message": "success",
		})
	}

}

func (u *ProductUsecase) CreateProd(c *gin.Context) {
	var prod entities.ProductRequest

	fmt.Println(" log access {}", prod.Name)
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(400, gin.H{
			"data":    err.Error(),
			"message": "error pointer payload",
		})
		return
	}
	log.Println("data logging {} ", prod)
	if err := u.barangRepo.CreateData(&prod); err != nil {

		c.JSON(400, gin.H{
			"data":    prod,
			"error":   err,
			"product": "Sample Product",
		})
	} else {
		c.JSON(200, gin.H{
			"data":    prod,
			"product": "Data berhasil disimpan",
		})

	}
}

func (pr *ProductUsecase) UpdateData(c *gin.Context) {
	var payload entities.Product

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(400, gin.H{
		"payload": payload.Name,
		"data":    nil,
		"message": "succss",
	})

}

func (u *ProductUsecase) GetProductInfo(c *gin.Context) {
	var prod = []entities.Product{{
		ID:   1,
		Name: "Sample Product",
	},
		{
			ID:   1,
			Name: "Sample Product",
		}, {
			ID:   1,
			Name: "Sample Product",
		}, {
			ID:   1,
			Name: "Sample Product",
		}, {
			ID:   1,
			Name: "Sample Product",
		},
	}
	c.JSON(200, gin.H{
		"data":    prod,
		"product": "Sample Product",
	})

}
