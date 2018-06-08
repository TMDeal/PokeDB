package models

import sq "github.com/Masterminds/squirrel"

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

func Versions() sq.SelectBuilder {
	return sq.Select("*").From("versions")
}

func VersionGroups() sq.SelectBuilder {
	return sq.Select("*").From("version_groups")
}

func (v Version) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := sq.Select("*").From("version_groups").Where("id = ?", v.VersionGroupID)
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}

func (vg VersionGroup) Generation(f Finder) (*Generation, error) {
	var gen Generation
	query := sq.Select("*").From("generations").Where("id = ?", vg.GenerationID)
	if err := f.Find(&gen, query); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (vg VersionGroup) Regions(f Finder) ([]Region, error) {
	var rs []Region
	query := sq.Select("*").
		From("regions").
		Join("version_group_regions ON regions.id = version_group_regions.region_id").
		Where("version_group_id = ?", vg.ID)

	if err := f.FindAll(&rs, query); err != nil {
		return nil, err
	}

	return rs, nil
}

func (vg VersionGroup) Versions(f Finder) ([]Version, error) {
	var vs []Version
	query := sq.Select("*").From("versions").Where("version_group_id = ?", vg.ID)
	if err := f.FindAll(&vs, query); err != nil {
		return nil, err
	}

	return vs, nil
}
