package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

//Region represents a region entry in the database
type Region struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

//RegionFinder says how to find information for a region model
type RegionFinder interface {
	FindRegions(search interface{}) ([]*Region, error)
}

func (db DB) FindRegions(search interface{}) ([]*Region, error) {
	var rs []*Region
	var stmt *sqlx.Stmt
	var err error

	baseQuery := `
	select * from regions %s
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
		var r Region
		err := rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}

	return rs, nil
}
