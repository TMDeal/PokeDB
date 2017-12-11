package models

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gedex/inflector"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DB is the database connection for the application. It contains all the
//functions required to find information from the database, as defined by the
//Finder interface
type DB struct {
	conn *sqlx.DB
}

//Finder is an interface that defines how to get data about pokemon
type Finder interface {
	Find(model interface{}, conds Builder) error
	FindAll(models interface{}, conds Builder) error
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

func (db DB) Count(table string) (int, error) {
	var count int

	query := fmt.Sprintf(`select count(*) from %s`, table)
	err := db.conn.QueryRowx(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db DB) Find(model interface{}, conds Builder) error {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("model must be a pointer")
	}
	t = t.Elem()

	table := strings.ToLower(strcase.ToSnake(inflector.Pluralize(t.Name())))
	conditions, args := conds.ToSQL()

	query := fmt.Sprintf(`SELECT %s.* FROM %s %s`, table, table, conditions)
	query = db.conn.Rebind(query)

	err := db.conn.QueryRowx(query, args...).StructScan(model)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) FindAll(models interface{}, conds Builder) error {
	t := reflect.TypeOf(models)
	v := reflect.ValueOf(models).Elem()

	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("models must be a pointer")
	}
	t = t.Elem()

	if t.Kind() != reflect.Slice {
		return fmt.Errorf("models must be a slice")
	}
	t = t.Elem()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	table := strings.ToLower(strcase.ToSnake(inflector.Pluralize(t.Name())))
	conditions, args := conds.ToSQL()

	query := fmt.Sprintf(`SELECT %s.* FROM %s %s`, table, table, conditions)
	query = db.conn.Rebind(query)
	rows, err := db.conn.Queryx(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		r := reflect.New(t)
		err = rows.StructScan(r.Interface())
		if err != nil {
			return err
		}
		v.Set(reflect.Append(v, r))
	}

	return nil
}
