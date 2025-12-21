package usecases

import (
	"rianRestapp/config"
	"rianRestapp/entities"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SuplierUsecase struct {
	db *gorm.DB
}

func NewSuplierUsecase() *SuplierUsecase {
	db, _ := config.NewDB()
	return &SuplierUsecase{
		db: db,
	}
}

func (r *SuplierUsecase) IndexData(c *gin.Context) {

	var suplier []entities.Suplier
	err := r.db.Find(&suplier).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": err,
			"sss":   "ss",
		})
	}

	c.JSON(200, gin.H{
		"data": suplier,
		"sss":  "ss",
	})
	return
}

func (db *SuplierUsecase) Create(c *gin.Context) {

	var payload entities.Suplier
	err := c.ShouldBindBodyWithJSON(&payload)
	if err != nil {
		c.JSON(500, gin.H{
			"err":      err.Error(),
			"messages": "payload error",
		})
		return
	}
	if err := db.db.Create(&payload).Error; err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"messages": "data  berasil di simpna",
	})

}

func (rp *SuplierUsecase) UpdateSuplier(c *gin.Context) {
	id := c.Param("id")
	filrest, _ := strconv.Atoi(id)
	var suplier entities.Suplier
	if err := rp.db.Where("id", filrest).Save(&suplier).Error; err != nil {
		c.JSON(200, gin.H{
			"messages": "data  berasil di simpna",
		})

		return
	}
	c.JSON(200, gin.H{
		"messages": "data  berasil di simpna",
	})
	return

}
func (dt *SuplierUsecase) Delete(c *gin.Context) {
	id := c.Param("id")
	paramid, _ := strconv.Atoi(id)
	if err := dt.db.Delete("id", paramid).Error; err != nil {
		c.JSON(200, gin.H{
			"messages": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"messages": "data berhasil dihapus.",
	})
	return
}

func (dt *SuplierUsecase) ShowById(c *gin.Context) {
	var data entities.Suplier
	paramid := c.Param("id")
	id, _ := strconv.Atoi(paramid)
	if err := dt.db.Where("id", id).First(&data).Error; err != nil {
		c.JSON(200, gin.H{
			"messages": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": data,
	})
	return

}
