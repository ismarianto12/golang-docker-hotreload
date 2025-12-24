package usecases

import (
	"log"
	"regexp"
	"rianRestapp/config"
	"rianRestapp/entities"
	"rianRestapp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserUsecase struct {
	db *gorm.DB
}

func NewUserusecase() *UserUsecase {
	dbconfig, _ := config.NewDB()
	return &UserUsecase{
		db: dbconfig,
	}
}

func (db *UserUsecase) IndexData(c *gin.Context) {
	var userentity []entities.User
	varperpage := c.DefaultQuery("perpage", "10")
	search := c.DefaultQuery("search", "")
	isStringOnly := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if !isStringOnly(search) {
		utils.BuildResponse(userentity, 200, "search hanya boleh huruf", c)
		return
	}
	perpage, err := strconv.Atoi(varperpage)
	if err != nil {
		utils.BuildResponse(userentity, 200, "payload error acces", c)
		return
	}
	log.Printf("paramload %s", perpage)

	if err := db.db.Find(&userentity).Error; err != nil {
		utils.BuildResponse(err.Error(), 500, err.Error(), c)
		return
	}
	utils.BuildResponse(userentity, 200, "data berhasil dload "+varperpage, c)
	return

}
