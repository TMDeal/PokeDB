package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

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
	FindGenerations(search interface{}) ([]*Generation, error)
}

//GenerationSelfFinder is an interface that says how a generation should
//find its own data relationship information
type GenerationSelfFinder interface {
	RegionFinder
}

//Region is a getter function for a Generations region info
func (g Generation) Region() (*Region, error) {
	rs, err := g.retriever.FindRegions(g.RegionID)
	if err != nil {
		return nil, err
	}

	return rs[0], nil
}

func (db DB) FindGenerations(search interface{}) ([]*Generation, error) {
	var gens []*Generation
	var stmt *sqlx.Stmt
	var err error

	baseQuery := `
	select g.id, g.main_region_id as "region_id", g.identifier, g.name from
	generations as g %s
	`

	switch search.(type) {
	case int:
		stmt, err = db.session.Preparex(fmt.Sprintf(baseQuery, `
		where id = $1
		`))
	case string:
		stmt, err = db.session.Preparex(fmt.Sprintf(baseQuery, `
		where lower(name) like lower($1%)
		`))
	default:
		return nil, ErrInvalidSearch
	}

	rows, err := stmt.Queryx(search)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var gen Generation
		gen.retriever = db

		err := rows.StructScan(&gen)
		if err != nil {
			return nil, err
		}
		gens = append(gens, &gen)
	}

	return gens, nil
}
