package models

import sq "github.com/Masterminds/squirrel"

type FlavorText struct {
	VersionGroupID int64  `db:"version_group_id"`
	Text           string `db:"flavor_text"`
}

func (flav FlavorText) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := sq.Select("*").From("version_groups").Where(sq.Eq{"id": flav.VersionGroupID})
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}
