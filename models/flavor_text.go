package models

type FlavorText struct {
	VersionGroupID int64  `db:"version_group_id"`
	Text           string `db:"flavor_text"`
}

func (flav FlavorText) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := Select("*").From("version_groups").Where("id = ?", flav.VersionGroupID)
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}
