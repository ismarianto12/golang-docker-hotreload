package entities

type TypeBarang struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Type       string `json:"type" gorm:"type:varchar(40);"`
	UserID     uint   `json:"user_id" gorm:"not null;index"`
	LocationId string `json:"location_id" gorm:"type:varchar(20);"`
}

func (TypeBarang) TableName() string {
	return "type_barang"
}
