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

func NewProductUsecase(barangRepo *repositories.BarangRepo) *ProductUsecase {
	return &ProductUsecase{barangRepo: barangRepo}
}

func (u *ProductUsecase) CreateProd(c *gin.Context) {
	var prod entities.Product

	fmt.Println(" log access {}", prod.Name)
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(400, gin.H{
			"data":    err.Error(),
			"message": "error pointer payload",
		})
		return
	}
	// log.Default("data" );
	log.Println("data logging {} ", prod)

	if err := u.barangRepo.CreateData(&prod).Error; err != nil {
		c.JSON(200, gin.H{
			"data":    prod,
			"product": "Sample Product",
		})
	} else {
		c.JSON(400, gin.H{
			"data":    prod,
			"product": "Gagal simpan data",
		})

	}
	return

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
