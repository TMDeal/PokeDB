package models

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DB is the database connection for the application. It contains all the
//functions required to find information from the database, as defined by the
//Finder interface
type DB struct {
	session *sqlx.DB
}

//Finder is an interface that defines how to get data about pokemon
type Finder interface {
	GenerationFinder
	RegionFinder
	DamageClassFinder
}

const (
	user    = "pokedb"
	pass    = ""
	host    = "localhost"
	dbName  = "pokedex"
	sslMode = "disable"
)

//NewDB returns a new DB and an error if a failure occurs
func NewDB() (*DB, error) {
	info := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, pass, host, dbName, sslMode)
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

	return &DB{session: db}, nil
}

func (db DB) Close() error {
	return db.session.Close()
}
