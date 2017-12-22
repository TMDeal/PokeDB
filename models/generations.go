package models

type Generation struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	RegionID   int    `db:"main_region_id"`
}

func Generations() *SelectBuilder {
	return Select("*").From("generations")
}

func (g Generation) Region(f Finder) (*Region, error) {
	var r Region
	query := Select("*").From("regions").Where("id = ?", g.RegionID)
	if err := f.Find(&r, query); err != nil {
		return nil, err
	}

	return &r, nil
}
