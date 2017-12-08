package models

import "database/sql"

type Move struct {
	ID                   int64         `db:"id"`
	Identifier           string        `db:"identifier"`
	Name                 string        `db:"name"`
	GenerationID         int64         `db:"generation_id"`
	TypeID               int64         `db:"type_id"`
	TargetID             int64         `db:"target_id"`
	DamageClassID        int64         `db:"damage_class_id"`
	ContestTypeID        sql.NullInt64 `db:"contest_type_id"`
	ContestEffectID      sql.NullInt64 `db:"contest_effect_id"`
	SuperContestEffectID sql.NullInt64 `db:"super_contest_effect_id"`
	EffectID             sql.NullInt64 `db:"effect_id"`
	Power                sql.NullInt64 `db:"power"`
	PP                   sql.NullInt64 `db:"pp"`
	Accuracy             sql.NullInt64 `db:"accuracy"`
	Priority             int64         `db:"priority"`
	EffectChance         sql.NullInt64 `db:"effect_chance"`
}

func (m Move) Targets(f Finder) (*MoveTarget, error) {
	var t MoveTarget
	if err := f.Find(&t, "id = ?", m.TargetID); err != nil {
		return nil, err
	}

	return &t, nil
}

func (m Move) Type(f Finder) (*Type, error) {
	var t Type
	if err := f.Find(&t, "id = ?", m.TypeID); err != nil {
		return nil, err
	}

	return &t, nil
}

func (m Move) Generation(f Finder) (*Generation, error) {
	var gen Generation
	if err := f.Find(&gen, "id = ?", m.GenerationID); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (m Move) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	if err := f.Find(&dc, "id = ?", m.DamageClassID); err != nil {
		return nil, err
	}

	return &dc, nil
}
