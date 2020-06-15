package main

import (
	"fmt"
	"os"
)

// PostgresUsername is the username for the Potgres connection
var PostgresUsername string = os.Getenv("POSTGRES_USERNAME")

// PostgresPassword is the password for the Potgres connection
var PostgresPassword string = os.Getenv("POSTGRES_PASSWORD")

// PostgresDatabase is the name of the databasae to connect to for the Potgres connection
var PostgresDatabase string = os.Getenv("POSTGRES_DATABASE")

// PostgresHost is the host for the Potgres connection
var PostgresHost string = os.Getenv("POSTGRES_HOST")

// PostgresConnectionString is the connection string for the Postgres connection
var PostgresConnectionString string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", PostgresHost, PostgresUsername, PostgresDatabase, PostgresPassword)