package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PostgresConnection is the globally shared database connection
var PostgresConnection *gorm.DB
var err error

func main() {

	PostgresConnection, err = CreatePostgresDatabaseConnection()
	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/new", NewPoll).Methods("POST")
}
