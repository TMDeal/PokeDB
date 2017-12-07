package models

import (
	"github.com/gocraft/dbr"
)

type Move struct {
	ID                   int64         `db:"id"`
	Identifier           string        `db:"identifier"`
	Name                 string        `db:"name"`
	GenerationID         int64         `db:"generation_id"`
	TypeID               int64         `db:"type_id"`
	TargetID             int64         `db:"target_id"`
	DamageClassID        int64         `db:"damage_class_id"`
	ContestTypeID        dbr.NullInt64 `db:"contest_type_id"`
	ContestEffectID      dbr.NullInt64 `db:"contest_effect_id"`
	SuperContestEffectID dbr.NullInt64 `db:"super_contest_effect_id"`
	EffectID             dbr.NullInt64 `db:"effect_id"`
	Power                dbr.NullInt64 `db:"power"`
	PP                   dbr.NullInt64 `db:"pp"`
	Accuracy             dbr.NullInt64 `db:"accuracy"`
	Priority             int64         `db:"priority"`
	EffectChance         dbr.NullInt64 `db:"effect_chance"`
}

type MoveFinder interface {
	FindMoves(limit uint64, offset uint64) ([]*Move, error)
	FindMove(query string, values ...interface{}) (*Move, error)
}

func (m Move) Targets(mtf MoveTargetFinder) (*MoveTarget, error) {
	mt, err := mtf.FindMoveTarget("id = ?", m.TargetID)
	if err != nil {
		return nil, err
	}

	return mt, nil
}

func (m Move) Type(tf TypeFinder) (*Type, error) {
	t, err := tf.FindType("id = ?", m.TypeID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (m Move) Generation(gf GenerationFinder) (*Generation, error) {
	gen, err := gf.FindGeneration("id = ?", m.GenerationID)
	if err != nil {
		return nil, err
	}

	return gen, err
}

func (m Move) DamageClass(dcf DamageClassFinder) (*DamageClass, error) {
	dc, err := dcf.FindDamageClass("id = ?", m.DamageClassID)
	if err != nil {
		return nil, err
	}

	return dc, nil
}

func (db DB) FindMoves(limit uint64, offset uint64) ([]*Move, error) {
	var ms []*Move
	sess := db.Session()

	_, err := sess.Select("*").
		From("moves").
		Limit(limit).
		Offset(offset).
		Load(&ms)

	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (db DB) FindMove(query string, values ...interface{}) (*Move, error) {
	var m Move
	sess := db.Session()

	_, err := sess.Select("*").
		From("moves").
		Where(query, values...).
		Load(&m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
