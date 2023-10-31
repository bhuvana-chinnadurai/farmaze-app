package model

import "database/sql"

type OrderStatuses interface {
	GetAll() ([]OrderStatus, error)
}

type OrderStatusRepository struct {
	DB *sql.DB
}

func NewOrderStatusRepository(db *sql.DB) OrderStatuses {
	return &OrderStatusRepository{DB: db}
}

func (r *OrderStatusRepository) GetAll() ([]OrderStatus, error) {
	var statuses []OrderStatus

	rows, err := r.DB.Query("SELECT id, name FROM order_statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var status OrderStatus
		err := rows.Scan(&status.ID, &status.Name)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}
