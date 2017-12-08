package models

type Generation struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	RegionID   int    `db:"main_region_id"`
}

func (g Generation) Region(f Finder) (*Region, error) {
	var r Region
	if err := f.Find(&r, "id = ?", g.RegionID); err != nil {
		return nil, err
	}

	return &r, nil
}
