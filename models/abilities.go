package models

import sq "github.com/Masterminds/squirrel"

type Ability struct {
	ID           int64  `db:"id"`
	Identifier   string `db:"identifier"`
	Name         string `db:"name"`
	GenerationID int64  `db:"generation_id"`
	ShortEffect  string `db:"short_effect"`
	Effect       string `db:"effect"`
}

func Abilities() sq.SelectBuilder {
	return sq.Select("*").From("abilities")
}

func (a Ability) Generation(f Finder) (*Generation, error) {
	var gen Generation
	query := sq.Select("*").From("generations").Where("id = ?", a.GenerationID)
	if err := f.Find(&gen, query); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (a Ability) FlavorText(f Finder, vg int64) (*FlavorText, error) {
	var flav FlavorText
	query := sq.Select("version_group_id", "flavor_text").From("ability_flavor_text").
		Where(sq.Eq{"ability_id": a.ID, "version_group_id": vg})

	if err := f.Find(&flav, query); err != nil {
		return nil, err
	}

	return &flav, nil
}
