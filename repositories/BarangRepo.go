package repositories

import (
	"rianRestapp/entities"

	"gorm.io/gorm"
)

type BarangRepo struct {
	db *gorm.DB
}

func NewBarangRepo(db *gorm.DB) *BarangRepo {
	return &BarangRepo{db: db}
}

func (r *BarangRepo) CreateData(prod *entities.Product) error {
	if err := r.db.Save(&prod).Error; err != nil {
		return err
	}
	return nil

}

func (r *BarangRepo) GetDB() *gorm.DB {
	return r.db
}
func (r *BarangRepo) SetDB(db *gorm.DB) {
	r.db = db
}
func (r *BarangRepo) WithTx(tx *gorm.DB) *BarangRepo {
	return &BarangRepo{db: tx}
}
func (r *BarangRepo) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
func (r *BarangRepo) CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}
