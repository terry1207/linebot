package repository

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB
var err error

func InitDbContext() {
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
}
