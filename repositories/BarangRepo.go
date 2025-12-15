package repositories

import (
	"log"
	"rianRestapp/config"
	"rianRestapp/entities"

	"gorm.io/gorm"
)

type BarangRepo struct {
	db *gorm.DB
}

func NewBarangRepo() *BarangRepo {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal("❌ gagal create DB:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ gagal ambil sql.DB:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("❌ DB TIDAK TERKONEKSI:", err)
	}

	log.Println("✅ DB TERKONEKSI")

	return &BarangRepo{db: db}
}

func (r *BarangRepo) GetAllData() ([]entities.Product, error) {
	var data []entities.Product
	if err := r.db.Select("name").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *BarangRepo) Delete(id int) error {
	if err := r.db.Delete(&id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BarangRepo) FindById(id int) (*entities.Product, error) {
	var data entities.Product
	if err := r.db.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *BarangRepo) SaveData(prod *entities.Product) error {
	if err := r.db.Save(&prod).Error; err != nil {
		return err
	}
	return nil
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
