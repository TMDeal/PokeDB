package models

type FlavorText struct {
	VersionGroupID int64  `db:"version_group_id"`
	Text           string `db:"flavor_text"`
}

func (flav FlavorText) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	if err := f.Find(&vg, NewConditions().Where("id = ?", flav.VersionGroupID)); err != nil {
		return nil, err
	}

	return &vg, nil
}
