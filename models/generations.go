package models

//Generation represents a generation entry in the database
type Generation struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	RegionID   int    `db:"region_id"`
}

//GenerationFinder says how to find information for a Generation
type GenerationFinder interface {
	FindGenerations(search interface{}) ([]*Generation, error)
}

//Region is a getter function for a Generations region info
func (g Generation) Region(rf RegionFinder) (*Region, error) {
	rs, err := rf.FindRegions(g.RegionID)
	if err != nil {
		return nil, err
	}

	return rs[0], nil
}

func (db DB) FindGenerations(search interface{}) ([]*Generation, error) {
	var gens []*Generation

	baseQuery := `
	select g.id, g.main_region_id as "region_id", g.identifier, g.name from
	generations as g %s
	`

	rows, err := db.GetRows(baseQuery, search)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var gen Generation

		err := rows.StructScan(&gen)
		if err != nil {
			return nil, err
		}
		gens = append(gens, &gen)
	}

	return gens, nil
}
