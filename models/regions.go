package models

import "log"

type Region struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func (db DB) FindRegionByID(id int32) (*Region, error) {
	var region Region

	err := db.QueryRowx(`
	select r.*, rn.name from
	regions as r, region_names as rn
	where
	r.id = rn.region_id
	and rn.local_language_id = 9
	`, id).StructScan(&region)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &region, nil
}
