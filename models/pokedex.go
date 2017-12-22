package models

import "database/sql"

type Pokedex struct {
	ID           int64         `db:"id"`
	RegionID     sql.NullInt64 `db:"region_id"`
	Identifier   string        `db:"identifier"`
	IsMainSeries bool          `db:"is_main_series"`
	Description  string        `db:"description"`
}

func Pokedexes() *SelectBuilder {
	return Select("*").From("pokedexes")
}

func (p Pokedex) Region(f Finder) (*Region, error) {
	if !p.RegionID.Valid {
		return nil, nil
	}

	var r Region
	query := Select("*").From("regions").Where("id = ?", p.RegionID.Int64)
	if err := f.Find(&r, query); err != nil {
		return nil, err
	}

	return &r, nil
}

func (p Pokedex) VersionGroups(f Finder) ([]VersionGroup, error) {
	var vgs []VersionGroup
	query := Select("*").From("pokemon_version_groups as pvg").
		Join("version_groups as vg ON pvg.version_group_id = vg.id").
		Where("pokedex_id = ?", p.ID)

	if err := f.FindAll(&vgs, query); err != nil {
		return nil, err
	}

	return vgs, nil
}
