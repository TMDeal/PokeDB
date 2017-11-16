package models

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

const (
	user    = "pokedb"
	pass    = ""
	dbName  = "pokedb"
	sslMode = "disable"
)

func NewDB() (*DB, error) {
	info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, pass, dbName, sslMode)
	db, err := sqlx.Open("postgres", info)
	if err != nil {
		log.Printf("Error opening database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error establishing database connection")
		return nil, err
	}

	return &DB{db}, nil
}
