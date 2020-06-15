package main

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// CreatePostgresDatabaseConnection creates a database connection instance
func CreatePostgresDatabaseConnection() (*gorm.DB, error) {
	postgresconnection, err := gorm.Open("postgres", PostgresConnectionString)
	if err != nil {
		return nil, errors.New("[ POSTGRES CONNECTION FAILED ] ")
	}
	return postgresconnection, nil
}
