package db

import "log"

type Stat struct {
	ID            int    `db:"id"`
	Identifier    string `db:"identifier"`
	AltIdentifier string `db:"alt_identifier"`
}

func (db DB) FindStatsByID(id int) (*Stat, error) {
	var s Stat

	err := db.QueryRowx(`
	select id, identifier, alt_identifier from stats where id = $1
	`, id).StructScan(&s)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &s, nil
}
