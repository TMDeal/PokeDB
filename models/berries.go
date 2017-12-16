package models

type Berry struct {
	ID                int64 `db:"id"`
	ItemID            int64 `db:"item_id"`
	FirmnessID        int64 `db:"firmness_id"`
	NaturalGiftPower  int64 `db:"natural_gift_power"`
	NaturalGiftTypeID int64 `db:"natural_gift_type_id"`
	Size              int64 `db:"size"`
	MaxHarvest        int64 `db:"max_harvest"`
	GrowthTime        int64 `db:"growth_time"`
	SoilDryness       int64 `db:"soil_dryness"`
	Smoothness        int64 `db:"smoothness"`
}

type BerryFlavor struct {
	BerryID       int64 `db:"berry_id"`
	ContestTypeID int64 `db:"contest_type_id"`
	Flavor        int64 `db:"flavor"`
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

func (b Berry) Item(f Finder) (*Item, error) {
	var i Item
	query := Select("*").From("items").Where("id = ?", b.ItemID)
	if err := f.Find(&i, query); err != nil {
		return nil, err
	}

	return &i, nil
}

func (b Berry) Flavors(f Finder) ([]*BerryFlavor, error) {
	var bfs []*BerryFlavor
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
