package models

import "database/sql"

type Location struct {
	ID         int64  `db:"id"`
	RegionID   int64  `db:"region_id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

type LocationArea struct {
	ID         int64          `db:"id"`
	LocationID int64          `db:"location_id"`
	Identifier sql.NullString `db:"identifier"`
	Name       sql.NullString `db:"name"`
}

type LocationEncounterMethod struct {
	EncounterMethod
	VersionID int64 `db:"version_id"`
	Rate      int64 `db:"rate"`
}

func Locations() *SelectBuilder {
	return Select("*").From("locations")
}

func (l LocationArea) Encounters(f Finder) ([]LocationEncounterMethod, error) {
	var le []LocationEncounterMethod
	query := Select("*").From("location_area_encounter_rates as laer").
		Join("encounter_methods as em ON em.id = laer.encounter_method_id").
		Where("laer.location_area_id = ?", l.ID)

	if err := f.Find(&le, query); err != nil {
		return nil, err
	}

	return le, nil

}

func (l Location) Region(f Finder) (*Region, error) {
	var r Region
	query := Select("*").From("regions").Where("id = ?", l.RegionID)
	if err := f.Find(&r, query); err != nil {
		return nil, err
	}

	return &r, nil
}

func (l Location) Areas(f Finder) ([]LocationArea, error) {
	var la []LocationArea
	query := Select("*").From("location_areas").Where("location_id = ?", l.ID)
	if err := f.FindAll(&la, query); err != nil {
		return nil, err
	}

	return la, nil
}
