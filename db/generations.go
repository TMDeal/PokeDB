package db

import (
	"log"
)

type Generation struct {
	ID     int     `db:"id"`
	Name   string  `db:"name"`
	Region *Region `db:"region"`
}

func (db DB) FindGenerationByID(id int) (*Generation, error) {
	var gen Generation

	err := db.QueryRowx(`
	select g.id, g.name,
	r.id as "region.id", r.name as "region.name" from
	generations as g, regions as r
	where g.region_id = r.id and g.id = $1
	`, id).StructScan(&gen)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &gen, nil

}

func (db DB) FindGenerationByName(name string) (*Generation, error) {
	var gen Generation

	err := db.QueryRowx(`
	select g.id, g.name,
	r.id as "region.id", r.name as "region.name" from
	generations as g, regions as r
	where g.region_id = r.id and g.name = $1
	`, name).StructScan(&gen)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &gen, nil

}
