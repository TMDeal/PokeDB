package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type Stat struct {
	ID            int64         `db:"id"`
	Identifier    string        `db:"identifier"`
	Name          string        `db:"name"`
	BattleOnly    bool          `db:"is_battle_only"`
	GameIndex     sql.NullInt64 `db:"game_index"`
	DamageClassID sql.NullInt64 `db:"damage_class_id"`
}

func Stats() sq.SelectBuilder {
	return sq.Select("*").From("Stats")
}

func (s Stat) DamageClass(f Finder) (*DamageClass, error) {
	if !s.DamageClassID.Valid {
		return nil, nil
	}

	var dc DamageClass
	query := sq.Select("*").From("damage_classes").Where("id = ?", s.DamageClassID.Int64)
	if err := f.Find(&dc, query); err != nil {
		return nil, err
	}

	return &dc, nil
}
