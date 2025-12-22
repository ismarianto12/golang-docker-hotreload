package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	var payloaddata []entities.StockMovement
	if err := c.ShouldBindJSON(&payloaddata).Error; err != nil {
		utils.BuildResponse(err, 400, err(), c)
		return
	}
	const msg string = "data berhasid di ubah"
	utils.BuildResponse(nil, 400, msg, c)
	return

}
func (repo *StockMovement) DeleteData(c *gin.Context) {
	var payloaddata []entities.StockMovement
	const msg string = "data berhasid di hapus"

	id := c.Param("id")
	restid, err := strconv.Atoi(id)
	if err != nil {
		utils.BuildResponse(nil, 400, msg, c)
	}
	if err := repo.db.Delete(&payloaddata).Where("id", restid).Error; err != nil {
		utils.BuildResponse(nil, 400, err.Error(), c)
	}
	utils.BuildResponse(nil, 400, msg, c)
	return

}

func (repo *StockMovement) CallApi(c *gin.Context) {
	resp, err := http.Get("https://www.mncsekuritas.id/backendweb/singleslide")
	resp.Header.Set("Authorization", "sss")
	if err != nil {
		fmt.Print("log data call err")
	}
	body, err := ioutil.ReadAll(resp.Body)
	var converdata []Convertdata
	defer resp.Body.Close()

	var resbody = string(body)
	var marshaljson = json.Unmarshal([]byte(resbody), &converdata)
	if marshaljson != nil {
		utils.BuildResponse(marshaljson.Error(), 400, "gagal parsing data", c)
		return
	}
	utils.BuildResponse(converdata, 200, "success data", c)

}

type Convertdata struct {
	Title    string `json:"title"`
	Headline string `json:"headline"`
}
