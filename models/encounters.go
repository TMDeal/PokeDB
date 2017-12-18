package models

type Encounter struct {
	ID              int64 `db:"id"`
	VersionID       int64 `db:"version_id"`
	LocationAreaID  int64 `db:"location_area_id"`
	EncounterSlotID int64 `db:"encounter_slot_id"`
	PokemonID       int64 `db:"pokemon_id"`
	MinLevel        int64 `db:"min_level"`
	MaxLevel        int64 `db:"max_level"`
}

type EncounterSlot struct {
	ID                int64 `db:"id"`
	VersionGroupID    int64 `db:"version_group_id"`
	EncounterMethodID int64 `db:"encounter_method_id"`
	Slot              int64 `db:"slot"`
	Rarity            int64 `db:"rarity"`
}

type EncounterMethod struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	Ordering   int64  `db:"ordering"`
}

type EncounterCondition struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

type EncounterConditionValue struct {
	ID                   int64  `db:"id"`
	EncounterConditionID int64  `db:"encounter_condition_id"`
	IsDefault            bool   `db:"is_default"`
	identifier           string `db:"identifier"`
	Name                 string `db:"name"`
}

func (e Encounter) Condition(f Finder) (*EncounterCondition, error) {
	var ec EncounterCondition
	query := Select("*").From("encounter_slot_id as ecv").
		Join("encounter_condition_value_map as ecvm ON ecv.id = ecvm.encounter_condition_value_id").
		Where("encounter_id = ?", e.ID)

	if err := f.Find(&ec, query); err != nil {
		return nil, err
	}

	return &ec, nil
}

func (e Encounter) Version(f Finder) (*Version, error) {
	var v Version
	query := Select("*").From("versions").Where("id = ?", e.VersionID)
	if err := f.Find(&v, query); err != nil {
		return nil, err
	}

	return &v, nil
}

func (e Encounter) Area(f Finder) (*LocationArea, error) {
	var la LocationArea
	query := Select("*").From("location_areas").Where("id = ", e.LocationAreaID)
	if err := f.Find(&la, query); err != nil {
		return nil, err
	}

	return &la, nil
}

func (e Encounter) Slot(f Finder) (*EncounterSlot, error) {
	var es EncounterSlot
	query := Select("*").From("encounter_slots").Where("id = ?", e.EncounterSlotID)
	if err := f.Find(&es, query); err != nil {
		return nil, err
	}

	return &es, nil
}

func (e EncounterSlot) Method(f Finder) (*EncounterMethod, error) {
	var em EncounterMethod
	query := Select("*").From("encounter_methods").Where("id = ?", e.EncounterMethodID)
	if err := f.Find(&em, query); err != nil {
		return nil, err
	}

	return &em, nil
}

func (e EncounterSlot) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := Select("*").From("version_groups").Where("id = ?", e.VersionGroupID)
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}

func (e EncounterCondition) Values(f Finder) ([]*EncounterConditionValue, error) {
	var ecv []*EncounterConditionValue
	query := Select("*").From("encounter_condition_values").Where("encounter_condition_id = ?", e.ID)
	if err := f.FindAll(&ecv, query); err != nil {
		return nil, err
	}

	return ecv, nil
}
