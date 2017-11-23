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

//TypeFinder is an interface that says how to find a Type
type TypeFinder interface {
	FindTypes(limit uint64, offset uint64) ([]*Type, error)
	FindType(query string, value interface{}) (*Type, error)
}

//Generation gets the generation info for a Type
func (t Type) Generation(gf GenerationFinder) (*Generation, error) {
	gen, err := gf.FindGeneration("id = ?", t.GenerationID)
	if err != nil {
		return nil, err
	}

	return gen, nil
}

//DamageClass gets the damage class info for a type
func (t Type) DamageClass(df DamageClassFinder) (*DamageClass, error) {
	dc, err := df.FindDamageClass("id = ?", t.DamageClassID)
	if err != nil {
		return nil, err
	}

	return dc, err
}

func (db DB) FindType(query string, value interface{}) (*Type, error) {
	var t Type
	sess := db.Session()

	_, err := sess.Select("*").From("types").Where(query, value).Load(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (db DB) FindTypes(limit uint64, offset uint64) ([]*Type, error) {
	var ts []*Type
	sess := db.Session()

	_, err := sess.Select("*").From("types").Limit(limit).Offset(offset).Load(&ts)
	if err != nil {
		return nil, err
	}

	return ts, nil
}
