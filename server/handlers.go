package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewPoll will create a new poll in the app
func NewPoll(w http.ResponseWriter, r *http.Request) {
	
}