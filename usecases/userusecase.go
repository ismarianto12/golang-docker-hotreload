package usecases

import (
	"log"
	"regexp"
	"rianRestapp/config"
	"rianRestapp/entities"
	"rianRestapp/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func (grm *UserUsecase) Create(c *gin.Context) {
	var user entities.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BuildResponse(nil, 200, "err"+err.Error(), c)
	}
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		{
			if err != nil {
				utils.BuildResponse(nil, 500, err.Error(), c)
				return
			}
		}
		user.Password = string(hash)
	}
	if err := grm.db.Create(&user).Error; err != nil {
		utils.BuildResponse(nil, 200, err.Error(), c)
		return

	}
	utils.BuildResponse(nil, 200, "Berhasil simpan data", c)

}

func (grm *UserUsecase) Update(c *gin.Context) {
	var user entities.User

	id := c.Param("id")
	filterId, err := strconv.Atoi(id)

	if err != nil {
		utils.BuildResponse(nil, 400, "error id"+err.Error(), c)
	}

	if err := grm.db.First(&user, filterId).Error; err != nil {
		utils.BuildResponse(nil, 404, "User tidak ditemukan", c)
		return
	}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		{
			if err != nil {
				utils.BuildResponse(nil, 500, err.Error(), c)
				return
			}
		}
		user.Password = string(hash)
	}

	if err := c.ShouldBindBodyWithJSON(&user).Error; err != nil {
		utils.BuildResponse(nil, 200, "err"+err(), c)
	}
	if err := grm.db.Save(&user).Error; err != nil {
		utils.BuildResponse(nil, 200, err.Error(), c)
		return
	}
	utils.BuildResponse(nil, 200, "Berhasil update", c)

}

func (grm *UserUsecase) Delete(c *gin.Context) {
	var user entities.User

	id := c.Param("id")
	filterId, err := strconv.Atoi(id)
	if err != nil {
		utils.BuildResponse(nil, 400, "error id"+err.Error(), c)
	}
	if err := c.ShouldBindJSON(&user).Error; err != nil {
		utils.BuildResponse(nil, 200, "err"+err(), c)
	}
	if err := grm.db.Where("id", filterId).Delete(&user).Error; err != nil {
		utils.BuildResponse(nil, 200, err.Error(), c)

	}
	utils.BuildResponse(nil, 200, "Berhasil simpan data", c)

}

func (grm *UserUsecase) Show(c *gin.Context) {
	var user entities.User

	id := c.Param("id")
	filterId, err := strconv.Atoi(id)
	if err != nil {
		utils.BuildResponse(nil, 400, "error id "+err.Error(), c)
		return
	}
	if err := grm.db.Where("id", filterId).First(&user).Error; err != nil {
		utils.BuildResponse(nil, 200, err.Error(), c)
		return
	}
	utils.BuildResponse(user, 200, "Berhasil simpan data", c)

}
