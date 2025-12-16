package usecases

import (
	"log"
	"rianRestapp/entities"
	"rianRestapp/repositories"
	"strconv"

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
	if err := c.ShouldBindBodyWithJSON(&typeBarang); err != nil {
		log.Print("❌ gagal create DB:", err)
		c.JSON(200, gin.H{
			"data": "data",
		})
		return

	}
	if err := rp.localrepo.InsertData(&typeBarang); err != nil {
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

}
func (rp *CategReUsecaseDetail) UpdateData(c *gin.Context) {
	var typebarang entities.TypeBarang
	if err := c.ShouldBindBodyWithJSON(&typebarang); err != nil {
		c.JSON(400, gin.H{
			"msg":   err,
			"data ": nil,
		})
		return
	}
	id := c.Param("id")
	safeid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"msg":   err.Error(),
			"data ": nil,
		})
		return

	}
	if err := rp.localrepo.UpdateByID(safeid); err != nil {
		c.JSON(200, gin.H{
			"msg":   err,
			"data ": "gaga update ata",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   "data berhasil di update",
		"data ": typebarang.Type,
	})

}
