package models

//Type represents a Type in the database
type Type struct {
	ID            int    `db:"id"`
	Identifier    string `db:"identifier"`
	GenerationID  int    `db:"generation_id"`
	DamageClassID int    `db:"damage_class_id"`
	Name          string `db:"name"`
}

//TypeFinder is an interface that says how to find a Type
type TypeFinder interface {
	FindTypes(search interface{}) ([]*Type, error)
	FindType(search interface{}) (*Type, error)
}

//Generation gets the generation info for a Type
func (t Type) Generation(gf GenerationFinder) (*Generation, error) {
	gen, err := gf.FindGeneration(t.GenerationID)
	if err != nil {
		return nil, err
	}

	return gen, nil
}

//DamageClass gets the damage class info for a type
func (t Type) DamageClass(df DamageClassFinder) (*DamageClass, error) {
	dc, err := df.FindDamageClass(t.DamageClassID)
	if err != nil {
		return nil, err
	}

	return dc, err
}

func (db DB) FindType(search interface{}) (*Type, error) {
	var t Type

	row, err := db.Row(`
	select * from types %s
	`, search)
	if err != nil {
		return nil, err
	}

	err = row.StructScan(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (db DB) FindTypes(search interface{}) ([]*Type, error) {
	var ts []*Type

	baseQuery := `
	select * from types %s
	`

	rows, err := db.Rows(baseQuery, search)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t Type

		err := rows.StructScan(&t)
		if err != nil {
			return nil, err
		}
		ts = append(ts, &t)
	}

	return ts, nil
}
