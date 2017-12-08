package models

type Stat struct {
	ID            int64  `db:"id"`
	Identifier    string `db:"identifier"`
	Name          string `db:"name"`
	BattleOnly    bool   `db:"is_battle_only"`
	GameIndex     int    `db:"game_index"`
	DamageClassID int    `db:"damage_class_id"`
}

func (s Stat) DamageClass(f Finder) (*DamageClass, error) {
	var dc DamageClass
	if err := f.Find(&dc, "id = ?", s.DamageClassID); err != nil {
		return nil, err
	}

	return &dc, nil
}
