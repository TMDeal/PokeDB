package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (db DB) Row(query string, search interface{}) (*sqlx.Row, error) {
	switch search.(type) {
	case int:
		return db.session.QueryRowx(fmt.Sprintf(query, `
		where id = $1
		`), search), nil

	case string:
		search = fmt.Sprintf(`%s%%`, search)
		return db.session.QueryRowx(fmt.Sprintf(query, `
		where lower(name) like lower($1)
		`), search), nil

	case nil:
		return db.session.QueryRowx(fmt.Sprintf(query, ``)), nil

	default:
		return nil, ErrInvalidSearch
	}
}

func (db DB) Rows(query string, search interface{}) (*sqlx.Rows, error) {
	switch search.(type) {
	case int:
		return db.session.Queryx(fmt.Sprintf(query, `
		where id = $1
		`), search)

	case string:
		search = fmt.Sprintf(`%s%%`, search)
		return db.session.Queryx(fmt.Sprintf(query, `
		where lower(name) like lower($1)
		`), search)

	case nil:
		return db.session.Queryx(fmt.Sprintf(query, ``))

	default:
		return nil, ErrInvalidSearch
	}
}
