package models

import sq "github.com/Masterminds/squirrel"

type Generation struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	RegionID   int    `db:"main_region_id"`
}

func Generations() sq.SelectBuilder {
	return sq.Select("*").From("generations")
}

func (g Generation) Region(f Finder) (*Region, error) {
	var r Region
	query := sq.Select("*").From("regions").Where("id = ?", g.RegionID)
	if err := f.Find(&r, query); err != nil {
		return nil, err
	}

	return &r, nil
}
