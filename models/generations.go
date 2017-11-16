package models

import (
	"log"
)

type Generation struct {
	ID         int     `db:"id"`
	Identifier string  `db:"identifier"`
	Region     *Region `db:"region"`
}

func (db DB) FindGenerationByID(id int) (*Generation, error) {
	var gen Generation

	err := db.QueryRowx(`
	select g.id, g.identifier,
	r.id as "region.id", r.identifier as "region.identifier" from
	generations as g, regions as r
	where g.region_id = r.id and g.id = $1
	`, id).StructScan(&gen)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &gen, nil

}
