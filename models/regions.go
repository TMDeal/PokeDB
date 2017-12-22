package models

//Region represents a region entry in the database
type Region struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func Regions() *SelectBuilder {
	return Select("*").From("regions")
}
