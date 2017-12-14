package models

type Version struct {
	ID             int64  `db:"id"`
	VersionGroupID int64  `db:"version_group_id"`
	Identifier     string `db:"identifier"`
	Name           string `db:"name"`
}

type VersionGroup struct {
	ID           int64  `db:"id"`
	Identifier   string `db:"identifier"`
	GenerationID int64  `db:"generation_id"`
	Ordering     int64  `db:"ordering"`
}

func (v Version) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	if err := f.Find(&vg, NewConditions().Where("id = ?", v.VersionGroupID)); err != nil {
		return nil, err
	}

	return &vg, nil
}

func (vg VersionGroup) Generation(f Finder) (*Generation, error) {
	var gen Generation
	if err := f.Find(&gen, NewConditions().Where("id = ?", vg.GenerationID)); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (vg VersionGroup) Regions(f Finder) ([]*Region, error) {
	var rs []*Region
	conds := NewConditions().
		Join("version_group_regions ON regions.id = version_group_regions.region_id").
		Where("version_group_id = ?", vg.ID)

	if err := f.FindAll(&rs, conds); err != nil {
		return nil, err
	}

	return rs, nil
}

func (vg VersionGroup) Versions(f Finder) ([]*Version, error) {
	var vs []*Version
	if err := f.FindAll(&vs, NewConditions().Where("version_group_id = ?", vg.ID)); err != nil {
		return nil, err
	}

	return vs, nil
}
