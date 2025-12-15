package repositories

import (
	"log"
	"rianRestapp/config"
	"rianRestapp/entities"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo() *CategoryRepo {
	local, err := config.NewDB()
	if err != nil {
		log.Fatal("❌ gagal create DB:", err)

	}
	return &CategoryRepo{db: local}
}

func (gr *CategoryRepo) GetAllData() ([]entities.TypeBarang, error) {
	var typebarang []entities.TypeBarang
	if err := gr.db.Select("type", "user_id").Find(&typebarang).Error; err != nil {
		return nil, err
	}
	return typebarang, nil
}

func (gr *CategoryRepo) InsertData(typebarang *entities.TypeBarang) error {
	err := gr.db.Create(&typebarang).Error
	if err != nil {
		return err
	}
	return nil
}

func (gr *CategoryRepo) UpdateData(id int) error {
	var category entities.TypeBarang
	if err := gr.db.Where("id", id).Save(&category).Error; err != nil {
		return err
	}
	return nil
}
func (gr *CategoryRepo) ShowId(id int) error {
	var category entities.TypeBarang
	if err := gr.db.Where("id", id).Find(&category).Error; err != nil {
		return err
	}
	return nil
}
func (gr *CategoryRepo) UpdateaId(id int) error {
	var category entities.TypeBarang
	if err := gr.db.Where("id", id).Find(&category).Error; err != nil {
		return err
	}
	return nil
}

func (gr *CategoryRepo) DeleteId(id int) error {
	var category entities.TypeBarang
	if err := gr.db.Where("id", id).Find(&category).Error; err != nil {
		return err
	}
	return nil
}
