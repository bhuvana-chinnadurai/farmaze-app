package model

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Orders interface {
	Create(order *Order) error
	List() ([]*Order, error)
	GetByClientID(clientID uuid.UUID) ([]Order, error)
}

type OrderStatus struct {
	ID   uuid.UUID
	Name string
}

// OrderRepository represents the repository for orders.
type OrderRepository struct {
	dbConn        *sql.DB
	orderProducts OrderProducts
}

// NewOrderRepository creates a new OrderRepository.
func NewOrderRepository(dbConn *sql.DB, orderProducts OrderProducts) Orders {
	return &OrderRepository{
		dbConn:        dbConn,
		orderProducts: orderProducts,
	}
}

func (r *OrderRepository) Create(order *Order) error {
	// Assuming you have a database connection available as r.dbConn
	tx, err := r.dbConn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	fmt.Println("order.Status.ID", order.Status.ID())
	// Insert the order details into the orders table
	_, err = tx.Exec(`INSERT INTO orders (id, client_id, total_price, created_at, status_id) VALUES ($1, $2, $3, $4, $5)`, order.ID, order.ClientID, order.TotalPrice, order.CreatedAt, order.Status.String())
	if err != nil {
		return fmt.Errorf("error while inserting to orders: %s", err.Error())
	}

	// Insert the order products into the order_products table
	for _, product := range order.Products {
		_, err = tx.Exec(`INSERT INTO order_products (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`, order.ID, product.ProductID, product.Quantity, product.Price)
		if err != nil {
			return fmt.Errorf("error while inserting to order_products: %s", err.Error())
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// List returns a list of all orders.
func (o *OrderRepository) List() ([]*Order, error) {
	query := `SELECT id, client_id, total_price, created_at, status FROM orders;`

	rows, err := o.dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order

	for rows.Next() {
		order := &Order{}
		err := rows.Scan(&order.ID, &order.ClientID, &order.TotalPrice, &order.CreatedAt, &order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetByClientID(clientID uuid.UUID) ([]Order, error) {
	rows, err := r.dbConn.Query("SELECT * FROM orders WHERE client_id = $1", clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.ClientID, &order.TotalPrice, &order.CreatedAt, &order.Status)
		if err != nil {
			return nil, err
		}

		// Retrieve products for the order
		products, err := r.orderProducts.GetProductsForOrder(order.ID)
		if err != nil {
			return nil, err
		}
		order.Products = products

		orders = append(orders, order)
	}

	return orders, nil
}
