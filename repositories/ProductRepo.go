package repositories

import (
	"database/sql"
	"rianRestapp/entities"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(database *sql.DB) *ProductRepo {
	return &ProductRepo{db: database}
}

func (r *ProductRepo) GetAllProducts() (entities.Product, error) {
	// Placeholder implementation
	var products entities.Product
	return products, nil

}

func (r *ProductRepo) GetProductByID(id int) (string, error) {
	// Placeholder implementation
	product := "Product" + string(id)
	return product, nil
}
