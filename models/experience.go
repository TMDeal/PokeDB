package models

type Experience struct {
	GrowthRateID int64 `db:"growth_rate_id"`
	Level        int64 `db:"level"`
	XP           int64 `db:"experience"`
}

type GrowthRate struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	Formula    string `db:"formula"`
}

func (e Experience) Growth(f Finder) (*GrowthRate, error) {
	var gr GrowthRate
	query := Select("*").From("growth_rates").Where("id = ?", e.GrowthRateID)
	if err := f.Find(&gr, query); err != nil {
		return nil, err
	}

	return &gr, nil
}
