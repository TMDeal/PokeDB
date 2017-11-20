package models

//DamageClass represets the damage class of a move or type. The damage classes
//are status, physical, or special
type DamageClass struct {
	ID          int    `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

//DamageClassFinder is an interface the defines ways to find a DamageClass
type DamageClassFinder interface {
	FindDamageClasses(search interface{}) ([]*DamageClass, error)
}

func (db DB) FindDamageClasses(search interface{}) ([]*DamageClass, error) {
	var dcs []*DamageClass

	baseQuery := `
	select * from move_damage_classes %s
	`

	rows, err := db.GetRows(baseQuery, search)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dc DamageClass
		err := rows.StructScan(&dc)
		if err != nil {
			return nil, err
		}
		dcs = append(dcs, &dc)
	}

	return dcs, nil
}
