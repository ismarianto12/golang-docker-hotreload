package usecases

import (
	"rianRestapp/config"
	"rianRestapp/entities"
	"rianRestapp/utils"
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

func (r *SuplierUsecase) UpdateDataImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
	}
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}
	contentType := file.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		c.JSON(400, gin.H{
			"message": "only jpeg and png images are allowed",
		})
		return
	}

	dest := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, dest); err != nil {
		c.JSON(400, gin.H{
			"Error": err,
			"sss":   "ss",
		})

	}
	utils.BuildResponse(nil, 200, "file upload success fully", c)

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
	utils.BuildResponse(suplier, 200, "success load data", c)

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
	utils.BuildResponse(nil, 200, "data berhasil disimpan", c)

}

func (rp *SuplierUsecase) UpdateSuplier(c *gin.Context) {
	id := c.Param("id")
	filrest, _ := strconv.Atoi(id)
	var suplier entities.Suplier
	if err := rp.db.Where("id", filrest).Save(&suplier).Error; err != nil {
		utils.BuildResponse(nil, 400, "gagal simpan data "+err.Error(), c)

		return
	}
	utils.BuildResponse(nil, 200, "data berhasil disimpan", c)
	return

}
func (dt *SuplierUsecase) Delete(c *gin.Context) {
	id := c.Param("id")
	var data entities.Suplier

	paramid, _ := strconv.Atoi(id)
	if err := dt.db.Where("id", paramid).Delete(data).Error; err != nil {
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
