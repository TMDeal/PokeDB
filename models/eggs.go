package models

type EggGroup struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func EggGroups() *SelectBuilder {
	return Select("*").From("egg_groups")
}
