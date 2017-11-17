package models

import (
	"log"
)

type Generation struct {
	ID         int     `db:"id"`
	Identifier string  `db:"identifier"`
	Name       string  `db:"name"`
	Region     *Region `db:"region"`
}

func (db DB) FindGenerationByID(id int32) (*Generation, error) {
	var gen Generation

	err := db.QueryRowx(`
	select g.id, g.identifier, gn.name,
	r.id as "region.id", r.identifier as "region.identifier" from
	generations as g, generation_names as gn, regions as r
	where 
    g.main_region_id = r.id and g.id = $1
    and g.id = gn.generation_id
    and gn.local_language_id = 9
	`, id).StructScan(&gen)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &gen, nil

}
