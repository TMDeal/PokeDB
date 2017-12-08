package models

type DamageClass struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
