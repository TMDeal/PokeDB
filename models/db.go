package models

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

//DB is the database connection for the application. It contains all the
//functions required to find information from the database, as defined by the
//Finder interface
type DB struct {
	conn *dbr.Connection
}

//Finder is an interface that defines how to get data about pokemon
type Finder interface {
	RegionFinder
	GenerationFinder
	TypeFinder
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
	db, err := dbr.Open("postgres", info, nil)
	if err != nil {
		log.Printf("Error connecting to database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error establishing database connection")
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (db DB) Close() error {
	return db.conn.Close()
}

func (db DB) Session() *dbr.Session {
	return db.conn.NewSession(nil)
}
