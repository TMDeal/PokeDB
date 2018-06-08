package models

import sq "github.com/Masterminds/squirrel"

type ContestType struct {
	ID         int    `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
	Flavor     string `db:"flavor"`
	Color      string `db:"color"`
}

type ContestEffect struct {
	ID         int    `db:"id"`
	Appeal     int    `db:"appeal"`
	Jam        int    `db:"jam"`
	FlavorText string `db:"flavor_text"`
	Effect     string `db:"effect"`
}

type SuperContestEffect struct {
	ID         int    `db:"id"`
	Appeal     string `db:"appeal"`
	FlavorText string `db:"flavor_text"`
}

type Combo struct {
	FirstMoveID  int `db:"first_move_id"`
	SecondMoveID int `db:"second_move_id"`
}

type ContestCombo struct {
	Combo
}

type SuperContestCombo struct {
	Combo
}

func ContestEffects() sq.SelectBuilder {
	return sq.Select("*").From("contest_effects")
}

func ContestTypes() sq.SelectBuilder {
	return sq.Select("*").From("contest_types")
}

func SuperContestEffects() sq.SelectBuilder {
	return sq.Select("*").From("super_contest_effects")
}

func ContestCombos() sq.SelectBuilder {
	return sq.Select("*").From("contest_combos")
}

func SuperContestCombos() sq.SelectBuilder {
	return sq.Select("*").From("super_contest_combos")
}

func (c Combo) FirstMove(f Finder) (*Move, error) {
	var m Move
	query := sq.Select("*").From("moves").Where("id = ?", c.FirstMoveID)
	if err := f.Find(&m, query); err != nil {
		return nil, err
	}

	return &m, nil
}

func (c Combo) SecondMove(f Finder) (*Move, error) {
	var m Move
	query := sq.Select("*").From("moves").Where("id = ?", c.SecondMoveID)
	if err := f.Find(&m, query); err != nil {
		return nil, err
	}

	return &m, nil
}
