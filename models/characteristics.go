package models

type Characteristic struct {
	ID          int64  `db:"id"`
	StatID      int64  `db:"stat_id"`
	GeneModFive int64  `db:"gene_mod_5"`
	Text        string `db:"text"`
}

func Characteristics() *SelectBuilder {
	return Select("*").From("characteristics")
}

func (c Characteristic) Stat(f Finder) (*Stat, error) {
	var s Stat
	query := Select("*").From("stats").Where("id = ?", c.StatID)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}
