package models

type Stat struct {
	ID            int64  `db:"id"`
	Identifier    string `db:"identifier"`
	Name          string `db:"name"`
	BattleOnly    bool   `db:"is_battle_only"`
	GameIndex     int    `db:"game_index"`
	DamageClassID int    `db:"damage_class_id"`
}

type StatFinder interface {
	FindStats(limit uint64, offset uint64) ([]*Stat, error)
	FindStat(query string, value interface{}) (*Stat, error)
}

func (s Stat) DamageClass(dcf DamageClassFinder) (*DamageClass, error) {
	dc, err := dcf.FindDamageClass("id = ?", s.DamageClassID)
	if err != nil {
		return nil, err
	}

	return dc, nil
}

func (db DB) FindStat(query string, value interface{}) (*Stat, error) {
	var s Stat
	sess := db.Session()

	_, err := sess.Select("*").
		From("stats").
		Where(query, value).
		Limit(1).
		Load(&s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (db DB) FindStats(limit uint64, offset uint64) ([]*Stat, error) {
	var ss []*Stat
	sess := db.Session()

	_, err := sess.Select("*").
		From("stats").
		Limit(limit).
		Offset(offset).
		Load(&ss)

	if err != nil {
		return nil, err
	}

	return ss, nil
}
