package models

import sq "github.com/Masterminds/squirrel"

type Stat struct {
	ID            int64  `db:"id"`
	Identifier    string `db:"identifier"`
	Name          string `db:"name"`
	BattleOnly    bool   `db:"is_battle_only"`
	GameIndex     int    `db:"game_index"`
	DamageClassID int    `db:"damage_class_id"`
}

func Stats() sq.SelectBuilder {
	return sq.Select("*").From("Stats")
}

func (s Stat) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	query := sq.Select("*").From("damage_classes").Where("id = ?", s.DamageClassID)
	if err := f.Find(&dc, query); err != nil {
		return nil, err
	}

	return &dc, nil
}
