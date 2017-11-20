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
	FindGeneration(search interface{}) (*Generation, error)
}

//Region is a getter function for a Generations region info
func (g Generation) Region(rf RegionFinder) (*Region, error) {
	r, err := rf.FindRegion(g.RegionID)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db DB) FindGeneration(search interface{}) (*Generation, error) {
	var gen Generation

	row, err := db.Row(`
	select g.id, g.main_region_id as "region_id", g.identifier, g.name from
	generations as g %s
	`, search)
	if err != nil {
		return nil, err
	}

	err = row.StructScan(&gen)
	if err != nil {
		return nil, err
	}

	return &gen, nil
}

func (db DB) FindGenerations(search interface{}) ([]*Generation, error) {
	var gens []*Generation

	baseQuery := `
	select g.id, g.main_region_id as "region_id", g.identifier, g.name from
	generations as g %s
	`

	rows, err := db.Rows(baseQuery, search)
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
