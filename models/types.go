package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

//Type represents a Type in the database
type Type struct {
	retriever     TypeSelfFinder
	ID            int    `db:"id"`
	Identifier    string `db:"identifier"`
	GenerationID  int    `db:"generation_id"`
	DamageClassID int    `db:"damage_class_id"`
	Name          string `db:"name"`
}

//TypeFinder is an interface that says how to find a Type
type TypeFinder interface {
	FindTypes(search interface{}) ([]*Type, error)
}

//TypeSelfFinder is an interface that says how a Type should find its own data
//relationship information
type TypeSelfFinder interface {
	GenerationFinder
	DamageClassFinder
}

//Generation gets the generation info for a Type
func (t Type) Generation() (*Generation, error) {
	gens, err := t.retriever.FindGenerations(t.GenerationID)
	if err != nil {
		return nil, err
	}

	return gens[0], nil
}

//DamageClass gets the damage class info for a type
func (t Type) DamageClass() (*DamageClass, error) {
	dcs, err := t.retriever.FindDamageClasses(t.DamageClassID)
	if err != nil {
		return nil, err
	}

	return dcs[0], err
}

func (db DB) FindTypes(search interface{}) ([]*Type, error) {
	var ts []*Type
	var stmt *sqlx.Stmt
	var err error

	baseQuery := `
	select * from types %s
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
		var t Type
		t.retriever = db

		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		ts = append(ts, &t)
	}

	return ts, nil
}