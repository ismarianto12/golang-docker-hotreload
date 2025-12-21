package entities

type Suplier struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama  string `json:"name" gorm:"type:varchar(50);null"`
	Jalan string `json:"jalan" gorm:"type:varchar(50);null"`
}

func (Suplier) TableName() string {
	return "suplier"
}
