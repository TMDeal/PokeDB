package models

import "database/sql"

type Pokemon struct {
	ID             int64  `db:"id"`
	Identifier     string `db:"identifier"`
	SpeciesID      int64  `db:"species_id"`
	Height         int64  `db:"height"`
	Weight         int64  `db:"weight"`
	BaseExperience int64  `db:"base_experience"`
	Ordering       int64  `db:"ordering"`
	IsDefault      bool   `db:"is_default"`
}

type PokemonForm struct {
	ID                         int64          `db:"id"`
	Identifier                 string         `db:"identifier"`
	FormIdentifier             sql.NullString `db:"form_identifier"`
	PokemonID                  int64          `db:"pokemon_id"`
	IntroducedInVersionGroupID int64          `db:"introduced_in_version_group_id"`
	Default                    bool           `db:"is_default"`
	BattleOnly                 bool           `db:"is_battle_only"`
	Mega                       bool           `db:"is_mega"`
	FormOrder                  int64          `db:"form_order"`
	Ordering                   int64          `db:"ordering"`
	PokemonName                sql.NullString `db:"pokemon_name"`
	FormName                   sql.NullString `db:"form_name"`
}

type PokemonAbility struct {
	Ability
	Hidden bool  `db:"is_hidden"`
	Slot   int64 `db:"slot"`
}

type PokemonItem struct {
	Item
	VersionID int64 `db:"version_id"`
	Rarity    int64 `db:"rarity"`
}

type PokemonMoveMethod struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type PokemonMove struct {
	Move
	VersionGroupID      int64         `db:"version_group_id"`
	PokemonMoveMethodID int64         `db:"pokemon_move_method_id"`
	Level               int64         `db:"level"`
	Ordering            sql.NullInt64 `db:"ordering"`
}

type PokemonStat struct {
	Stat
	BaseStat int64 `db:"base_stat"`
	Effort   int64 `db:"effort"`
}

type PokemonType struct {
	Type
	Slot int64 `db:"slot"`
}

func Pokemons() *SelectBuilder {
	return Select("*").From("pokemon")
}

func (p Pokemon) Forms(f Finder) ([]PokemonForm, error) {
	var pf []PokemonForm
	query := Select("*").From("pokemon_forms").Where("pokemon_id = ?", p.ID)
	if err := f.FindAll(&pf, query); err != nil {
		return nil, err
	}

	return pf, nil
}

func (p Pokemon) Species(f Finder) (*PokemonSpecies, error) {
	var s PokemonSpecies
	query := Select("*").From("pokemon_species").Where("id = ?", p.SpeciesID)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (p Pokemon) Abilities(f Finder) ([]PokemonAbility, error) {
	var a []PokemonAbility
	query := Select("a.*", "pa.is_hidden", "pa.slot").From("pokemon_abilities AS pa").
		Join("abilities AS a ON a.id = pa.ability_id").
		Where("pokemon_id = ?", p.ID)

	if err := f.FindAll(&a, query); err != nil {
		return nil, err
	}

	return a, nil
}

func (p Pokemon) Items(f Finder) ([]PokemonItem, error) {
	var i []PokemonItem
	query := Select("i.*", "pi.rarity", "pi.version_id").From("pokemon_items AS pi").
		Join("items AS i ON i.id = pi.item_id").
		Where("pokemon_id = ?", p.ID)

	if err := f.FindAll(&i, query); err != nil {
		return nil, err
	}

	return i, nil
}

func (p Pokemon) Stats(f Finder) ([]PokemonStat, error) {
	var s []PokemonStat
	query := Select("s.*", "ps.base_stat", "ps.effort").From("pokemon_stats AS ps").
		Join("stats AS s ON s.id = ps.stat_id").
		Where("pokemon_id = ?", p.ID)

	if err := f.FindAll(&s, query); err != nil {
		return nil, err
	}

	return s, nil
}

func (p Pokemon) Moves(f Finder) ([]PokemonMove, error) {
	var m []PokemonMove
	query := Select("m.*", "pm.level", "pm.ordering", "pm.pokemon_move_method_id", "pm.version_group_id").
		From("pokemon_moves AS pm").
		Join("moves AS m ON m.id = pm.move_id").
		Where("pokemon_id = ?", p.ID)

	if err := f.FindAll(&m, query); err != nil {
		return nil, err
	}

	return m, nil
}

func (p Pokemon) Types(f Finder) ([]PokemonType, error) {
	var t []PokemonType
	query := Select("t.*", "pt.slot").From("pokemon_types AS pt").
		Join("types AS t ON t.id = pt.type_id").
		Where("pokemon_id = ?", p.ID)

	if err := f.FindAll(&t, query); err != nil {
		return nil, err
	}

	return t, nil
}

func (pf PokemonForm) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := Select("*").From("version_groups").Where("id = ?", pf.IntroducedInVersionGroupID)
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}

func (pi PokemonItem) Version(f Finder) (*Version, error) {
	var v Version
	query := Select("*").From("versions").Where("id = ?", pi.VersionID)
	if err := f.Find(&v, query); err != nil {
		return nil, err
	}

	return &v, nil
}

func (pm PokemonMove) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := Select("*").From("version_groups").Where("id = ?", pm.VersionGroupID)
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}

func (pm PokemonMove) MoveMethod(f Finder) (*PokemonMoveMethod, error) {
	var m PokemonMoveMethod
	query := Select("*").From("pokemon_move_methods").Where("id = ?", pm.PokemonMoveMethodID)
	if err := f.Find(&m, query); err != nil {
		return nil, err
	}

	return &m, nil
}
