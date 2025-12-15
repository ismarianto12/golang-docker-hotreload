package entities

import "time"

type Penjualan struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Harga     string    `json:"name" gorm:"type:varchar(200);not null"`
	Stock     string    `json:"name" gorm:"type:varchar(200);not null"`
	HargaJual string    `json:"harga_jual" gorm:"type:varchar(100);not null"`
	HargaBeli string    `json:"harga_beli" gorm:"type:varchar(100);not null"`
	BarangId  string    `json:"barang_id" gorm:"type:varchar(200);not null"`
	UserId    string    `json:"user_id" gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `json:"updated_at" gorm:"index"; null`
	UpdatedAt time.Time `json:"updated_at" gorm:"index"; null`
}

func (Penjualan) TableName() string {
	return "penjualan"
}
