package model

import (
	"database/sql"
)

type Products interface {
	GetAllProducts() ([]Product, error)
}

type ProductRepository struct {
	dbConn *sql.DB
}

func NewProductRepository(dbConn *sql.DB) Products {
	return &ProductRepository{
		dbConn: dbConn,
	}
}

func (pr *ProductRepository) GetAllProducts() ([]Product, error) {
	rows, err := pr.dbConn.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.AvailableQuantity, &product.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
