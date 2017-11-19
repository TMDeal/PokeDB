package models

import "log"

//Generation represents a generation entry in the database
type Generation struct {
	retriever  GenerationSelfFinder
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	RegionID   int    `db:"region_id"`
}

//GenerationFinder says how to find information for a Generation
type GenerationFinder interface {
	FindGenerationByID(id int) (*Generation, error)
}

//GenerationSelfFinder is an interface that says how a generation should
//find its own data relationship information
type GenerationSelfFinder interface {
	RegionFinder
}

//Region is a getter function for a Generations region info
func (g Generation) Region() (*Region, error) {
	return g.retriever.FindRegionByID(g.RegionID)
}

//FindGenerationByID returns a Generation from the database based on the
//id of the entry. An error is also returned if a failure occured
func (db DB) FindGenerationByID(id int) (*Generation, error) {
	var gen Generation
	gen.retriever = db

	err := db.session.QueryRowx(`
	select g.id, g.main_region_id as "region_id", g.identifier, g.name from
	generations as g where g.id = $1
	`, id).StructScan(&gen)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &gen, nil

}
