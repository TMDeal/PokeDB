package models

//Region represents a region entry in the database
type Region struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

//RegionFinder says how to find information for a region model
type RegionFinder interface {
	FindRegions(search interface{}) ([]*Region, error)
	FindRegion(search interface{}) (*Region, error)
}

func (db DB) FindRegion(search interface{}) (*Region, error) {
	var r Region

	row, err := db.GetRow(`
	select * from regions %s
	`, search)
	if err != nil {
		return nil, err
	}

	row.StructScan(&r)

	return &r, nil
}

func (db DB) FindRegions(search interface{}) ([]*Region, error) {
	var rs []*Region

	baseQuery := `
	select * from regions %s
	`

	rows, err := db.GetRows(baseQuery, search)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r Region
		err := rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}

	return rs, nil
}
