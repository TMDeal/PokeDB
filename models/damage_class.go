package models

import "github.com/gocraft/dbr"

//DamageClass represets the damage class of a move or type. The damage classes
//are status, physical, or special
type DamageClass struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

//DamageClassFinder is an interface the defines ways to find a DamageClass
type DamageClassFinder interface {
	FindDamageClasses(limit uint64, offset uint64) ([]*DamageClass, error)
	FindDamageClass(query string, value interface{}) (*DamageClass, error)
}

func (db DB) FindDamageClass(query string, value interface{}) (*DamageClass, error) {
	var dc DamageClass
	sess := db.Session()

	count, err := sess.Select("*").From("move_damage_classes").Where(query, value).Load(&dc)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &dc, nil
}

func (db DB) FindDamageClasses(limit uint64, offset uint64) ([]*DamageClass, error) {
	var dcs []*DamageClass
	sess := db.Session()

	count, err := sess.Select("*").From("move_damage_classes").Limit(limit).Offset(offset).Load(&dcs)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return dcs, nil
}
