package entities

type TypeBarang struct {
	Id int `json:id,gorm:"primaryKey;autoIncrement"`
}

func TableName() string {
	return "type_barang"
}
