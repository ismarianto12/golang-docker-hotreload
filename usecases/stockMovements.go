package usecases

import (
	"log"
	"rianRestapp/config"
	"rianRestapp/entities"
	"rianRestapp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StockMovement struct {
	db *gorm.DB
}

func NewStockMovement() *StockMovement {
	db, _ := config.NewDB()
	return &StockMovement{db: db}
}
func (repo *StockMovement) IndeXalldata(c *gin.Context) {
	var alldata []entities.StockMovement
	log.Printf("detail get index movement {}")
	if err := repo.db.Find(&alldata).Error; err != nil {
		utils.BuildResponse(nil, 200, err.Error(), c)
		return
	}
	utils.BuildResponse(alldata, 200, utils.SUCCESS, c)
	return
}

func (repo *StockMovement) CreatedData(c *gin.Context) {
	var payloadmovement entities.StockMovement

	if err := c.ShouldBindBodyWithJSON(&payloadmovement).Error; err != nil {
		utils.BuildResponse(nil, 400, err(), c)
	}
	if err := repo.db.Create(&payloadmovement).Error; err != nil {
		utils.BuildResponse(nil, 400, "error "+err.Error(), c)
	}
	utils.BuildResponse(nil, 400, "data berhasil di simpan", c)
}

func (repo *StockMovement) ShowDatat(c *gin.Context) {
	id := c.Param("id")
	resultId, _ := strconv.Atoi(id)
	var payloadmovement entities.StockMovement
	if err := c.ShouldBindBodyWithJSON(&payloadmovement).Error; err != nil {
		utils.BuildResponse(nil, 400, "Mesage payload", c)
	}
	if err := repo.db.Where("id", resultId).Scan(&payloadmovement).Error; err != nil {
		utils.BuildResponse(nil, 400, "Error load data"+err.Error(), c)

	}
	utils.BuildResponse(nil, 200, "Data berahsil disimpan", c)
}
func (repo *StockMovement) UpdateData(c *gin.Context) {

}
func (repo *StockMovement) DeleteData(c *gin.Context) {

}
