package model

import (
	"database/sql"
	"fmt"
)

// Users defines the behaviors required for user operations.
type Users interface {
	Save(user *User) error
	// You can add more methods here as needed, such as FindByID, Delete, etc.
}

// UserRepository handles the operations with the database for user entities.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository with the given database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save inserts the User instance into the database.
// It ensures UserRepository implements the Users interface.
func (repo *UserRepository) Save(user *User) error {
	query := `INSERT INTO users (username, password, role) VALUES (?, ?, ?)`

	_, err := repo.db.Exec(query, user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %w", err)
	}
	return nil
}

// FindByUsername retrieves a user by their username.
func (repo *UserRepository) FindByUsername(username string) (*User, error) {
	var user User
	query := `SELECT id, username, password, role FROM users WHERE username = ?`
	err := repo.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
