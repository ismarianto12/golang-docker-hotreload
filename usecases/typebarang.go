package usecases

import (
	"log"
	"rianRestapp/entities"
	"rianRestapp/repositories"

	"github.com/gin-gonic/gin"
)

type CategReUsecaseDetail struct {
	localrepo *repositories.CategoryRepo
}

func NewTypeBarangUses() *CategReUsecaseDetail {
	typebarangRepo := repositories.NewCategoryRepo()
	return &CategReUsecaseDetail{localrepo: typebarangRepo}
}

func (rp *CategReUsecaseDetail) Alldata(c *gin.Context) {
	var entities entities.TypeBarang
	data, err := rp.localrepo.GetAllData()
	if err != nil {
		c.JSON(400, gin.H{
			"data": "data",
			"msg":  "errr",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": data, "message": "successfully handler data of",
		"code":   200,
		"entity": entities,
	})
}

func (rp *CategReUsecaseDetail) InserData(c *gin.Context) {
	var typeBarang entities.TypeBarang
	if err := c.ShouldBindBodyWithJSON(&typeBarang).Error; err != nil {
		log.Print("❌ gagal create DB:", err)
		c.JSON(200, gin.H{
			"data": "data",
		})
		return

	}
	if err := rp.localrepo.InsertData(&typeBarang).Error; err != nil {
		log.Print("❌ gagal create DB:", err)
		c.JSON(200, gin.H{
			"error":   err,
			"data":    "data",
			"message": "error message",
		})
		return

	}
	c.JSON(200, gin.H{
		"data": "data",
	})
	return

}
