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
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type MoveTarget struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type MoveFlavorText struct {
	FlavorText
	MoveID int64 `db:"move_id"`
}

type MoveEffect struct {
	ID          int64  `db:"id"`
	ShortEffect string `db:"short_effect"`
	Effect      string `db:"effect"`
}

func (m Move) Effect(f Finder) (*MoveEffect, error) {
	var me MoveEffect
	query := Select("*").From("move_effects").Where("id = ?", m.EffectID)
	if err := f.Find(&me, query); err != nil {
		return nil, err
	}

	return &me, nil
}

func (m Move) FlavorText(f Finder, vg int) (*MoveFlavorText, error) {
	var mft MoveFlavorText
	query := Select("*").From("move_flavor_text").
		Where("move_id = ?", m.ID).
		And("version_group_id = ?", vg)

	if err := f.Find(&mft, query); err != nil {
		return nil, err
	}

	return &mft, nil
}

func (m Move) Target(f Finder) (*MoveTarget, error) {
	var mt MoveTarget
	query := Select("*").From("move_targets").Where("id = ?", m.TargetID)
	if err := f.Find(&mt, query); err != nil {
		return nil, err
	}

	return &mt, nil
}

func (m Move) Flags(f Finder) ([]*MoveFlag, error) {
	var mf []*MoveFlag
	query := Select("*").
		From("move_flags").
		Join("move_flag_map ON move_flags.id = move_flag_map.move_flag_id").
		Where("move_flag_map.move_id = ?", m.ID)

	if err := f.FindAll(&mf, query); err != nil {
		return nil, err
	}

	return mf, nil
}

func (m Move) SuperContestEffect(f Finder) (*SuperContestEffect, error) {
	if !m.SuperContestEffectID.Valid {
		return nil, nil
	}

	var sce SuperContestEffect
	query := Select("*").From("super_contest_effects").Where("id = ?", m.SuperContestEffectID)
	if err := f.Find(&sce, query); err != nil {
		return nil, err
	}

	return &sce, nil
}

func (m Move) ContestType(f Finder) (*ContestType, error) {
	if !m.ContestTypeID.Valid {
		return nil, nil
	}

	var ct ContestType
	query := Select("*").From("contest_types").Where("id = ?", m.ContestTypeID.Int64)
	if err := f.Find(&ct, query); err != nil {
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
	query := Select("*").From("contest_effects").Where("id = ?", m.ContestEffectID.Int64)
	if err := f.Find(&ce, query); err != nil {
		return nil, err
	}

	return &ce, nil
}

func (m Move) Type(f Finder) (*Type, error) {
	var t Type
	query := Select("*").From("types").Where("id = ?", m.TypeID)
	if err := f.Find(&t, query); err != nil {
		return nil, err
	}

	return &t, nil
}

func (m Move) Generation(f Finder) (*Generation, error) {
	var gen Generation
	query := Select("*").From("generations").Where("id = ?", m.GenerationID)
	if err := f.Find(&gen, query); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (m Move) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	query := Select("*").From("damage_classes").Where("id = ?", m.DamageClassID)
	if err := f.Find(&dc, query); err != nil {
		return nil, err
	}

	return &dc, nil
}
