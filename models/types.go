package models

import "github.com/gocraft/dbr"

//Type represents a Type in the database
type Type struct {
	ID            int64         `db:"id"`
	Identifier    string        `db:"identifier"`
	GenerationID  int           `db:"generation_id"`
	DamageClassID dbr.NullInt64 `db:"damage_class_id"`
	Name          string        `db:"name"`
}

//Generation gets the generation info for a Type
func (t Type) Generation(f Finder) (*Generation, error) {
	var gen Generation
	if err := f.Find(&gen, "id = ?", t.GenerationID); err != nil {
		return nil, err
	}

	return &gen, nil
}

//DamageClass gets the damage class info for a type
func (t Type) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	if err := f.Find(&dc, "id = ?", t.DamageClassID); err != nil {
		return nil, err
	}

	return &dc, nil
}
