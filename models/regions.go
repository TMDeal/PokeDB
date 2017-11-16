package models

import "log"

type Region struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
}

func (db DB) FindRegionByID(id int) (*Region, error) {
	var region Region

	err := db.QueryRowx(`
	select id, identifier from regions where id = $1
	`, id).StructScan(&region)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &region, nil
}
