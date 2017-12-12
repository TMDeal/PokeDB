package models

import "database/sql"

type MoveMeta struct {
	MoveID         int64         `db:"move_id"`
	MetaCategoryID int64         `db:"meta_category_id"`
	MetaAilmentID  int64         `db:"meta_ailment_id"`
	MinHits        sql.NullInt64 `db:"min_hits"`
	MaxHits        sql.NullInt64 `db:"max_hits"`
	MinTurns       sql.NullInt64 `db:"min_turns"`
	MaxTurns       sql.NullInt64 `db:"max_turns"`
	Drain          int64         `db:"drain"`
	Healing        int64         `db:"healing"`
	CritRate       int64         `db:"crit_rate"`
	AilmentChance  int64         `db:"ailment_chance"`
	FlinchChance   int64         `db:"flinch_chance"`
	StatChance     int64         `db:"stat_chance"`
}

type MoveMetaAilment struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

type MoveMetaCategory struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Description string `db:"description"`
}

type MoveMetaStatChange struct {
	MoveID int64 `db:"move_id"`
	StatID int64 `db:"stat_id"`
	Change int64 `db:"change"`
}

func (m Move) Meta(f Finder) (*MoveMeta, error) {
	var meta MoveMeta
	if err := f.Find(&meta, NewConditions().Where("move_id = ?", m.ID)); err != nil {
		return nil, err
	}

	return &meta, nil
}

func (m MoveMeta) Category(f Finder) (*MoveMetaCategory, error) {
	var cat MoveMetaCategory
	if err := f.Find(&cat, NewConditions().Where("id = ?", m.MetaCategoryID)); err != nil {
		return nil, err
	}

	return &cat, nil
}

func (m MoveMeta) Ailment(f Finder) (*MoveMetaAilment, error) {
	var ail MoveMetaAilment
	if err := f.Find(&ail, NewConditions().Where("id = ?", m.MetaAilmentID)); err != nil {
		return nil, err
	}

	return &ail, nil
}

func (m MoveMeta) StatChanges(f Finder) ([]*MoveMetaStatChange, error) {
	var msc []*MoveMetaStatChange
	if err := f.FindAll(&msc, NewConditions().Where("move_id = ?", m.MoveID)); err != nil {
		return nil, err
	}

	return msc, nil
}

func (m MoveMetaStatChange) Stat(f Finder) (*Stat, error) {
	var s Stat
	if err := f.Find(&s, NewConditions().Where("id = ?", m.StatID)); err != nil {
		return nil, err
	}

	return &s, nil
}
