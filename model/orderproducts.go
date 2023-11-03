package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type OrderProducts interface {
	GetProductsForOrder(orderID uuid.UUID) ([]OrderProduct, error)
}

type OrderProductRepository struct {
	dbConn *sql.DB
}

func NewOrderProductRepository(dbConn *sql.DB) OrderProducts {
	return &OrderProductRepository{
		dbConn: dbConn,
	}
}

func (o *OrderProductRepository) GetProductsForOrder(orderID uuid.UUID) ([]OrderProduct, error) {
	rows, err := o.dbConn.Query("SELECT order_id, product_id, quantity FROM order_products WHERE order_id = $1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderProducts []OrderProduct
	for rows.Next() {
		var orderProduct OrderProduct
		err := rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)
		if err != nil {
			return nil, err
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}
