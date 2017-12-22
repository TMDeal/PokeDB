package models

type GrowthRate struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	Formula    string `db:"formula"`
}

type Experience struct {
	Level int64 `db:"level"`
	XP    int64 `db:"experience"`
}

func GrowthRates() *SelectBuilder {
	return Select("*").From("growth_rates")
}

func (e GrowthRate) Experience(f Finder) ([]Experience, error) {
	var gr []Experience
	query := Select("level", "experience").From("experience").Where("growth_rate_id = ?", e.ID)
	if err := f.FindAll(&gr, query); err != nil {
		return nil, err
	}

	return gr, nil
}
