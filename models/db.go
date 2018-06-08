package models

import (
	"fmt"
	"log"
	"reflect"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DB is the database connection for the application. It contains all the
//functions required to find information from the database, as defined by the
//Finder interface
type DB struct {
	conn   *sqlx.DB
	logger *log.Logger
}

//Finder is an interface that defines how to get data about pokemon
type Finder interface {
	Find(dest interface{}, conds sq.Sqlizer) error
	FindAll(dest interface{}, conds sq.Sqlizer) error
}

const (
	user    = "pokedb"
	pass    = ""
	host    = "localhost"
	dbName  = "pokedex"
	sslMode = "disable"
)

//NewDB returns a new DB and an error if a failure occurs
func NewDB(logger *log.Logger) (*DB, error) {
	info := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, pass, host, dbName, sslMode)

	db, err := sqlx.Open("postgres", info)
	if err != nil {
		logger.Printf("Error connecting to database")
		return nil, err
	}
	logger.Println("Connected to database successfully!")

	err = db.Ping()
	if err != nil {
		logger.Printf("Error establishing database connection")
		return nil, err
	}
	logger.Println("Connection to database established!")

	return &DB{conn: db, logger: logger}, nil
}

func (db DB) Log(msg ...interface{}) {
	db.logger.Println(msg)
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

func (db DB) Find(dest interface{}, stmt sq.Sqlizer) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("model must be a pointer")
	}

	sql, args, err := stmt.ToSql()
	if err != nil {
		return err
	}
	sql = db.conn.Rebind(sql)

	err = db.conn.QueryRowx(sql, args...).StructScan(dest)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) FindAll(dest interface{}, stmt sq.Sqlizer) error {
	t := reflect.TypeOf(dest)
	v := reflect.ValueOf(dest).Elem()

	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("dest must be a pointer")
	}
	t = t.Elem()

	if t.Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a slice")
	}
	t = t.Elem()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	sql, args, err := stmt.ToSql()
	if err != nil {
		return err
	}
	sql = db.conn.Rebind(sql)

	rows, err := db.conn.Queryx(sql, args...)
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
		v.Set(reflect.Append(v, r.Elem()))
	}

	return nil
}
