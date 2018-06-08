package models

import sq "github.com/Masterminds/squirrel"

type EggGroup struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func EggGroups() sq.SelectBuilder {
	return sq.Select("*").From("egg_groups")
}
