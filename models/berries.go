package models

type Berry struct {
	Item              `db:"item"`
	ID                int64 `db:"id"`
	FirmnessID        int64 `db:"firmness_id"`
	NaturalGiftPower  int64 `db:"natural_gift_power"`
	NaturalGiftTypeID int64 `db:"natural_gift_type_id"`
	Size              int64 `db:"size"`
	MaxHarvest        int64 `db:"max_harvest"`
	GrowthTime        int64 `db:"growth_time"`
	SoilDryness       int64 `db:"soil_dryness"`
	Smoothness        int64 `db:"smoothness"`
}

func Berries() *SelectBuilder {
	cols := []string{
		`items.id as "item.id"`,
		`items.identifier as "item.identifier"`,
		`items.category_id as "item.category_id"`,
		`items.cost as "item.cost"`,
		`items.fling_power as "item.fling_power"`,
		`items.fling_effect_id as "item.fling_effect_id"`,
		`items.name as "item.name"`,
		`items.short_effect as "item.short_effect"`,
		`items.effect as "item.effect"`,
		"berries.id",
		"berries.firmness_id",
		"berries.natural_gift_power",
		"berries.natural_gift_type_id",
		"berries.size",
		"berries.max_harvest",
		"berries.growth_time",
		"berries.soil_dryness",
		"berries.smoothness",
	}

	return Select(cols...).From("berries").
		Join("items ON items.id = berries.item_id")
}

type BerryFlavor struct {
	Name          string `db:"flavor"`
	Potency       int64  `db:"potency"`
	BerryID       int64  `db:"berry_id"`
	ContestTypeID int64  `db:"contest_type_id"`
}

type BerryFirmness struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func (b Berry) NaturalGiftType(f Finder) (*Type, error) {
	var t Type
	query := Select("*").From("types").Where("id = ?", b.NaturalGiftTypeID)
	if err := f.Find(&t, query); err != nil {
		return nil, err
	}

	return &t, nil
}

func (b Berry) Firmness(f Finder) (*BerryFirmness, error) {
	var bf BerryFirmness
	query := Select("*").From("berry_firmness").Where("id = ?", b.FirmnessID)
	if err := f.Find(&bf, query); err != nil {
		return nil, err
	}

	return &bf, nil
}

func (b Berry) Flavors(f Finder) ([]BerryFlavor, error) {
	var bfs []BerryFlavor
	query := Select("*").From("berry_flavors").Where("berry_id = ?", b.ID)

	if err := f.FindAll(&bfs, query); err != nil {
		return nil, err
	}

	return bfs, nil
}

func (b BerryFlavor) ContestType(f Finder) (*ContestType, error) {
	var ct ContestType
	query := Select("*").From("contest_types").Where("id = ?", b.ContestTypeID)
	if err := f.Find(&ct, query); err != nil {
		return nil, err
	}

	return &ct, nil
}
