package db

import "log"

type Region struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (db DB) FindRegionByID(id int) (*Region, error) {
	var region Region

	err := db.QueryRowx(`
	select id, name from regions where id = $1
	`, id).StructScan(&region)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &region, nil
}

func (db DB) FindRegionByName(name string) (*Region, error) {
	var region Region

	err := db.QueryRowx(`
	select id, name from regions where name = $1
	`, name).StructScan(&region)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &region, nil
}
