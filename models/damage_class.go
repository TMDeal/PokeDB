package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

//DamageClass represets the damage class of a move or type. The damage classes
//are status, physical, or special
type DamageClass struct {
	ID          int    `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

//DamageClassFinder is an interface the defines ways to find a DamageClass
type DamageClassFinder interface {
	FindDamageClasses(search interface{}) ([]*DamageClass, error)
}

func (db DB) FindDamageClasses(search interface{}) ([]*DamageClass, error) {
	var dcs []*DamageClass
	var stmt *sqlx.Stmt
	var err error

	baseQuery := `
	select * from move_damage_classes %s
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
		var dc DamageClass
		err := rows.StructScan(&dc)
		if err != nil {
			return nil, err
		}
		dcs = append(dcs, &dc)
	}

	return dcs, nil
}
