package models

import sq "github.com/Masterminds/squirrel"

//Region represents a region entry in the database
type Region struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func Regions() sq.SelectBuilder {
	return sq.Select("*").From("regions")
}
