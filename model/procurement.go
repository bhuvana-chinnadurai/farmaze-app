package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Procurements interface {
	Create(products []OrderProduct, createdAt time.Time) error
	ListByDate(createdAt time.Time) ([]*OrderProduct, error)
}

type ProcurementRepository struct {
	dbConn *sql.DB
}

func NewProcurementRepository(db *sql.DB) Procurements {
	return &ProcurementRepository{
		dbConn: db,
	}
}

func (p *ProcurementRepository) Create(products []OrderProduct, createdAt time.Time) error {
	tx, err := p.dbConn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	fmt.Println("products", products, createdAt)
	stmt, err := tx.Prepare("INSERT INTO procurement (product_id, quantity, created_at) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	defer stmt.Close()

	for _, product := range products {
		fmt.Println(product.ProductID, product.Quantity, createdAt.String())
		_, err := stmt.Exec(product.ProductID, product.Quantity, createdAt)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (p *ProcurementRepository) ListByDate(createdAt time.Time) ([]*OrderProduct, error) {
	fmt.Println(createdAt, createdAt)
	query := `SELECT product_id, quantity FROM procurement WHERE DATE(created_at) = $1`

	rows, err := p.dbConn.Query(query, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var procurements []*OrderProduct

	for rows.Next() {
		procurement := &OrderProduct{}
		err := rows.Scan(&procurement.ProductID, &procurement.Quantity)
		if err != nil {
			return nil, err
		}
		procurements = append(procurements, procurement)
	}

	return procurements, nil
}
