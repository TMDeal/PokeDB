package db

import "log"

type Stat struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	AltName string `db:"alt_name"`
}

func (db DB) FindStatsByID(id int) (*Stat, error) {
	var s Stat

	err := db.QueryRowx(`
	select id, name, alt_name from stats where id = $1
	`, id).StructScan(&s)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &s, nil
}
