package models

import "log"

//Region represents a region entry in the database
type Region struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

//RegionFinder says how to find information for a region model
type RegionFinder interface {
	FindRegionByID(id int) (*Region, error)
}

//FindRegionByID returns a Region from the database based on the id of the
//entry. An error is also returned if a failure occured
func (db DB) FindRegionByID(id int) (*Region, error) {
	var r Region

	err := db.session.QueryRowx(`
	select * from regions where id = $1
	`, id).StructScan(&r)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &r, nil
}
