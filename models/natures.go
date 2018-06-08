package models

import (
	sq "github.com/Masterminds/squirrel"
)

type Nature struct {
	ID              int64  `db:"id"`
	Identifier      string `db:"identifier"`
	DecreasedStatID int64  `db:"decreased_stat_id"`
	IncreasedStatID int64  `db:"increased_stat_id"`
	HatesFlavorID   int64  `db:"hates_flavor_id"`
	LikesFlavorID   int64  `db:"likes_flavor_id"`
	GameIndex       int64  `db:"game_index"`
	Name            string `db:"name"`
}

func Natures() sq.SelectBuilder {
	return sq.Select("*").From("natures")
}

func (n Nature) Decreased(f Finder) (*Stat, error) {
	var s Stat
	query := sq.Select("*").From("stats").Where("id = ?", n.DecreasedStatID)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (n Nature) Increased(f Finder) (*Stat, error) {
	var s Stat
	query := sq.Select("*").From("stats").Where("id = ?", n.IncreasedStatID)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (n Nature) Likes(f Finder) (*ContestType, error) {
	var ct ContestType
	query := sq.Select("*").From("contest_types").Where("id = ?", n.LikesFlavorID)
	if err := f.Find(&ct, query); err != nil {
		return nil, err
	}

	return &ct, nil
}

func (n Nature) Hates(f Finder) (*ContestType, error) {
	var ct ContestType
	query := sq.Select("*").From("contest_types").Where("id = ?", n.HatesFlavorID)
	if err := f.Find(&ct, query); err != nil {
		return nil, err
	}

	return &ct, nil
}
