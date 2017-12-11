package models

import (
	"database/sql"
	"log"
)

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

type MoveFlag struct {
	ID          int64
	Identifier  string
	Name        string
	Description string
}

type MoveTarget struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type MoveFlavorText struct {
	MoveID         int64  `db:"move_id"`
	VersionGroupID int64  `db:"version_group_id"`
	Text           string `db:"flavor_text"`
}

func (m Move) FlavorText(f Finder, vg int) (*MoveFlavorText, error) {
	var mft MoveFlavorText
	if err := f.Find(&mft, NewConditions().Where("move_id = ?", m.ID).And("version_group_id = ?", vg)); err != nil {
		return nil, err
	}

	return &mft, nil
}

func (m Move) Target(f Finder) (*MoveTarget, error) {
	var mt MoveTarget
	if err := f.Find(&mt, NewConditions().Where("id = ?", m.TargetID)); err != nil {
		return nil, err
	}

	return &mt, nil
}

func (m Move) Flags(f Finder) ([]*MoveFlag, error) {
	var mf []*MoveFlag
	conds := NewConditions().
		Join("move_flag_map ON move_flags.id = move_flag_map.move_flag_id").
		Where("move_flag_map.move_id = ?", m.ID)

	if err := f.FindAll(&mf, conds); err != nil {
		return nil, err
	}

	return mf, nil
}

func (m Move) SuperContestEffect(f Finder) (*SuperContestEffect, error) {
	if !m.SuperContestEffectID.Valid {
		return nil, nil
	}

	var sce SuperContestEffect
	if err := f.Find(&sce, NewConditions().Where("id = ?", m.SuperContestEffectID.Int64)); err != nil {
		return nil, err
	}

	return &sce, nil
}

func (m Move) ContestType(f Finder) (*ContestType, error) {
	if !m.ContestTypeID.Valid {
		return nil, nil
	}

	var ct ContestType
	if err := f.Find(&ct, NewConditions().Where("id = ?", m.ContestTypeID.Int64)); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ct, nil
}

func (m Move) ContestEffect(f Finder) (*ContestEffect, error) {
	if !m.ContestTypeID.Valid {
		return nil, nil
	}

	var ce ContestEffect
	if err := f.Find(&ce, NewConditions().Where("id = ?", m.ContestEffectID.Int64)); err != nil {
		return nil, err
	}

	return &ce, nil
}

func (m Move) Type(f Finder) (*Type, error) {
	var t Type
	if err := f.Find(&t, NewConditions().Where("id = ?", m.TypeID)); err != nil {
		return nil, err
	}

	return &t, nil
}

func (m Move) Generation(f Finder) (*Generation, error) {
	var gen Generation
	if err := f.Find(&gen, NewConditions().Where("id = ?", m.GenerationID)); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (m Move) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	if err := f.Find(&dc, NewConditions().Where("id = ?", m.DamageClassID)); err != nil {
		return nil, err
	}

	return &dc, nil
}
