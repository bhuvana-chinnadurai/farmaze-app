package model

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type Products interface {
	GetAll() ([]Product, error)
	GetByID(uuid.UUID) (*Product, error)
	Create(product *Product) error
	Edit(product *Product) error
	Delete(id uuid.UUID) error
}

type ProductRepository struct {
	dbConn *sql.DB
}

func NewProductRepository(dbConn *sql.DB) Products {
	return &ProductRepository{
		dbConn: dbConn,
	}
}

func (pr *ProductRepository) GetAll() ([]Product, error) {
	rows, err := pr.dbConn.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.AvailableQuantity, &product.Category, &product.Unit)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *ProductRepository) GetByID(productID uuid.UUID) (*Product, error) {
	product := &Product{}
	err := p.dbConn.QueryRow("SELECT * FROM products WHERE id = $1", productID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.AvailableQuantity,
		&product.Category,
		&product.Unit,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError{Message: "Product not found"}
		}
		log.Println("Error retrieving product:", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) Create(product *Product) error {
	_, err := p.dbConn.Exec("INSERT INTO products (id, name, price, description, available_quantity, category, unit) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		product.ID, product.Name, product.Price, product.Description, product.AvailableQuantity, product.Category, product.Unit)

	if err != nil {
		log.Println("Error creating product:", err)
		return err
	}

	return nil
}

func (p *ProductRepository) Edit(product *Product) error {
	_, err := p.dbConn.Exec("UPDATE products SET name=$2, price=$3, description=$4, available_quantity=$5, category=$6, unit=$7 WHERE id=$1",
		product.ID, product.Name, product.Price, product.Description, product.AvailableQuantity, product.Category, product.Unit)

	if err != nil {
		log.Println("Error editing product:", err)
		return err
	}

	return nil
}

func (p *ProductRepository) Delete(id uuid.UUID) error {
	_, err := p.dbConn.Exec("DELETE FROM products WHERE id=$1", id)

	if err != nil {
		log.Println("Error deleting product:", err)
		return err
	}

	return nil
}
