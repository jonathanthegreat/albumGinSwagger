package config

import (
	"github.com/go-pg/pg/v9"
	"log"
	"os"
)

// Connect to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "db_username",
		Password: "db_password",
		Addr:     "localhost:5432",
		Database: "db_name",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
