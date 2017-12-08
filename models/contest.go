package models

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

type ContestCombo struct {
	FirstMoveID  int `db:"first_move_id"`
	SecondMoveID int `db:"second_move_id"`
}

type SuperContestEffect struct {
	ID         int    `db:"id"`
	Appeal     string `db:"appeal"`
	FlavorText string `db:"flavor_text"`
}

type SuperContestCombo struct {
	FirstMoveID  int `db:"first_move_id"`
	SecondMoveID int `db:"second_move_id"`
}
