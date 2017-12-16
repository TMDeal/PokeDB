package models

import "database/sql"

type Item struct {
	ID            int64         `db:"id"`
	Identifier    string        `db:"identifier"`
	Name          string        `db:"name"`
	ShortEffect   string        `db:"short_effect"`
	Effect        string        `db:"effect"`
	Cost          int64         `db:"cost"`
	FlingPower    sql.NullInt64 `db:"fling_power"`
	FlingEffectID sql.NullInt64 `db:"fling_effect_id"`
	CategoryID    int64         `db:"category_id"`
}

type ItemFlavorText struct {
	FlavorText
	ItemID int64 `db:"item_id"`
}

type ItemCategory struct {
	ID         int64  `db:"id"`
	PocketID   int64  `db:"pocket_id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

type ItemFlag struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type ItemFlingEffect struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Effect     string `db:"effect"`
}

type ItemPocket struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func (i Item) FlavorText(f Finder, vg int) (*ItemFlavorText, error) {
	var flav ItemFlavorText
	query := Select("*").From("item_flavor_text").
		Where("item_id = ?", i.ID).
		And("version_group_id = ?", vg)

	if err := f.Find(&flav, query); err != nil {
		return nil, err
	}

	return &flav, nil
}

func (i Item) Flags(f Finder) ([]*ItemFlag, error) {
	var flags []*ItemFlag
	query := Select("*").
		From("item_flags").
		Join("item_flag_map ON item_flags.id = item_flag_map.item_flag_id").
		Where("item_flag_map.item_id = ?", i.ID)

	if err := f.FindAll(&flags, query); err != nil {
		return nil, err
	}

	return flags, nil
}

func (i Item) FlingEffect(f Finder) (*ItemFlingEffect, error) {
	if !i.FlingEffectID.Valid {
		return nil, nil
	}

	var e ItemFlingEffect
	query := Select("*").From("item_fling_effects").Where("id = ?", i.FlingEffectID.Int64)
	if err := f.Find(&e, query); err != nil {
		return nil, err
	}

	return &e, nil
}

func (i Item) Category(f Finder) (*ItemCategory, error) {
	var c ItemCategory
	query := Select("*").From("item_categories").Where("id = ?", i.CategoryID)
	if err := f.Find(&c, query); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c ItemCategory) Pocket(f Finder) (*ItemPocket, error) {
	var p ItemPocket
	query := Select("*").From("item_pockets").Where("id = ?", c.PocketID)
	if err := f.Find(&p, query); err != nil {
		return nil, err
	}

	return &p, nil
}
