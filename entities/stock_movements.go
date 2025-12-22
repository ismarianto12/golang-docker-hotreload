package entities

type StockMovement struct {
	ID            uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductId     int     `json:"product_id"`
	WhareHoiseId  int     `json:"wharehouse_id"`
	LocationFrom  int     `json:"location_from"`
	LocationTo    int     `json:"location_to"`
	Qty           float64 `json:"qty"`
	Type          int     `json:"type"`
	ReferenceType int     `json:"reference_type"`
	ReferenceId   int     `json:"reference_id"`
	CreateDat     int     `json:"created_at"`
	UpdateDat     int     `json:"updated_at"`
}

func (StockMovement) TableName() string {
	return "StockMovement"
}
