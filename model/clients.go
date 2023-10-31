package model

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type B2BClients interface {
	GetAllClients() ([]B2BClient, error)
	GetClientByID(id uuid.UUID) (*B2BClient, error)
}

type B2BClientRepository struct {
	dbConn *sql.DB
}

func NewB2BClientRepository(dbConn *sql.DB) B2BClients {
	return &B2BClientRepository{
		dbConn: dbConn,
	}
}

func (b *B2BClientRepository) GetAllClients() ([]B2BClient, error) {
	rows, err := b.dbConn.Query("SELECT id, company_name, contact_name, email, phone_number FROM b2b_clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []B2BClient
	for rows.Next() {
		var client B2BClient
		err := rows.Scan(&client.ID, &client.CompanyName, &client.ContactName, &client.Email, &client.PhoneNumber)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (c *B2BClientRepository) GetClientByID(clientID uuid.UUID) (*B2BClient, error) {
	client := &B2BClient{}
	err := c.dbConn.QueryRow("SELECT * FROM b2b_clients WHERE id = $1", clientID).Scan(
		&client.ID,
		&client.CompanyName,
		&client.ContactName,
		&client.Email,
		&client.PhoneNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError{Message: "Client not found"}
		}
		log.Println("Error retrieving client:", err)
		return nil, err
	}

	return client, nil
}
