package models

import "github.com/gocraft/dbr"

//Generation represents a generation entry in the database
type Generation struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	RegionID   int    `db:"region_id"`
}

//GenerationFinder says how to find information for a Generation
type GenerationFinder interface {
	FindGenerations(limit uint64) ([]*Generation, error)
	FindGeneration(query string, value interface{}) (*Generation, error)
}

//Region is a getter function for a Generations region info
func (g Generation) Region(rf RegionFinder) (*Region, error) {
	r, err := rf.FindRegion("id = ?", g.RegionID)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db DB) FindGeneration(query string, value interface{}) (*Generation, error) {
	var gen Generation
	sess := db.Session()

	count, err := sess.Select(
		"id", "main_region_id as region_id",
		"identifier", "name").From("generations").Where(query, value).Load(&gen)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &gen, nil
}

func (db DB) FindGenerations(limit uint64) ([]*Generation, error) {
	var gens []*Generation
	sess := db.Session()

	count, err := sess.Select(
		"id", "main_region_id as region_id",
		"identifier", "name").From("generations").Limit(limit).Load(&gens)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return gens, nil
}
