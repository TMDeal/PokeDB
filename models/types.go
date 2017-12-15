package models

import "database/sql"

type DamageMult int

const (
	DoubleDamage DamageMult = 200
	NormalDamage DamageMult = 100
	HalfDamage   DamageMult = 50
	NoDamage     DamageMult = 0
)

//Type represents a Type in the database
type Type struct {
	ID            int64         `db:"id"`
	Identifier    string        `db:"identifier"`
	GenerationID  int           `db:"generation_id"`
	DamageClassID sql.NullInt64 `db:"damage_class_id"`
	Name          string        `db:"name"`
}

func (t Type) DamageTo(f Finder, fact DamageMult) ([]*Type, error) {
	var ts []*Type
	conds := NewConditions().
		Join("type_efficacies as te on types.id = te.target_type_id").
		Where("te.damage_type_id = ?", t.ID).
		And("te.damage_factor = ?", fact)

	if err := f.FindAll(&ts, conds); err != nil {
		return nil, err
	}

	return ts, nil
}

func (t Type) DamageFrom(f Finder, fact DamageMult) ([]*Type, error) {
	var ts []*Type
	conds := NewConditions().
		Join("type_efficacies as te on types.id = te.damage_type_id").
		Where("te.target_type_id = ?", t.ID).
		And("te.damage_factor = ?", fact)

	if err := f.FindAll(&ts, conds); err != nil {
		return nil, err
	}

	return ts, nil
}

//Generation gets the generation info for a Type
func (t Type) Generation(f Finder) (*Generation, error) {
	var gen Generation
	if err := f.Find(&gen, NewConditions().Where("id = ?", t.GenerationID)); err != nil {
		return nil, err
	}

	return &gen, nil
}

//DamageClass gets the damage class info for a type
func (t Type) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	if err := f.Find(&dc, NewConditions().Where("id = ?", t.DamageClassID)); err != nil {
		return nil, err
	}

	return &dc, nil
}
