package models

import sq "github.com/Masterminds/squirrel"

type DamageClass struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func DamageClasses() sq.SelectBuilder {
	return sq.Select("*").From("damage_classes")
}
