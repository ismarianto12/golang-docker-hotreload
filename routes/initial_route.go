package routes

import "rianRestapp/usecases"

type Usecases struct {
	Product       *usecases.ProductUsecase
	TypeBarang    *usecases.CategReUsecaseDetail
	StockMovement *usecases.StockMovement
	Suplier       *usecases.SuplierUsecase
}

func NewUsecases() *Usecases {
	return &Usecases{
		Product:       usecases.NewProductUsecase(),
		TypeBarang:    usecases.NewTypeBarangUses(),
		StockMovement: usecases.NewStockMovement(),
		Suplier:       usecases.NewSuplierUsecase(),
	}
}
