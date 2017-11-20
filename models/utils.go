package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (db DB) GetRows(query string, search interface{}) (*sqlx.Rows, error) {
	var rows *sqlx.Rows
	var err error

	switch search.(type) {
	case int:
		rows, err = db.session.Queryx(fmt.Sprintf(query, `
		where id = $1
		`), search)
		if err != nil {
			return nil, err
		}

	case string:
		search = fmt.Sprintf(`%s%%`, search)
		rows, err = db.session.Queryx(fmt.Sprintf(query, `
		where lower(name) like lower($1)
		`), search)
		if err != nil {
			return nil, err
		}

	case nil:
		rows, err = db.session.Queryx(fmt.Sprintf(query, ``))
		if err != nil {
			return nil, err
		}

	default:
		return nil, ErrInvalidSearch
	}

	return rows, nil
}
