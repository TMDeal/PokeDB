package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type EvolutionChain struct {
	SpeciesID            int64         `db:"species_id"`
	EvolutionDetailsID   sql.NullInt64 `db:"evolution_details_id"`
	IsBaby               bool          `db:"is_baby"`
	BabyTriggerItemID    sql.NullInt64 `db:"baby_trigger_item_id"`
	EvolvesToSpeciesID   sql.NullInt64 `db:"evolves_to_species_id"`
	EvolvesFromSpeciesID sql.NullInt64 `db:"evolves_from_species_id"`
}

type EvolutionDetail struct {
	ID                    int64          `db:"id"`
	EvolvedSpeciesID      int64          `db:"evolved_species_id"`
	EvolutionTriggerID    int64          `db:"evolution_trigger_id"`
	TriggerItemID         sql.NullInt64  `db:"trigger_item_id"`
	MinimumLevel          int64          `db:"minimum_level"`
	Gender                sql.NullString `db:"gender"`
	Time                  sql.NullString `db:"time"`
	LocationID            sql.NullInt64  `db:"location_id"`
	HeldItemID            sql.NullInt64  `db:"held_item_id"`
	KnownMoveID           sql.NullInt64  `db:"known_move_id"`
	KnownMoveTypeID       sql.NullInt64  `db:"known_move_type_id"`
	MinimumHappiness      sql.NullInt64  `db:"minimum_happiness"`
	MinimumBeauty         sql.NullInt64  `db:"minimum_beauty"`
	MinimumAffection      sql.NullInt64  `db:"minimum_affection"`
	RelativePhysicalStats sql.NullInt64  `db:"relative_physical_stats"`
	PartySpeciesID        sql.NullInt64  `db:"party_species_id"`
	PartyTypeID           sql.NullInt64  `db:"party_type_id"`
	TradeSpeciesID        sql.NullInt64  `db:"trade_species_id"`
	NeedsOverworldRain    bool           `db:"needs_overworld_rain"`
	TurnUpsideDown        bool           `db:"turn_upside_down"`
}

type EvolutionTrigger struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

func EvolutionChains() sq.SelectBuilder {
	cols := []string{
		"ps.id as species_id",
		"pe.id as evolution_details_id",
		"ps.is_baby",
		"ec.baby_trigger_item_id",
		"ps.evolves_from_species_id",
		"(select id from pokemon_species where evolves_from_species_id = ps.id) as evolves_to_species_id",
	}

	return sq.Select(cols...).From("pokemon_species as ps").
		Join("evolution_chains AS ec ON ps.evolution_chain_id = ec.id").
		LeftJoin("pokemon_evolution AS pe ON pe.evolved_species_id = ps.id")
}

func (e EvolutionDetail) Trigger(f Finder) (*EvolutionTrigger, error) {
	var t EvolutionTrigger
	query := sq.Select("*").From("evolution_triggers").Where("id = ?", e.EvolutionTriggerID)
	if err := f.Find(&t, query); err != nil {
		return nil, err
	}

	return &t, nil
}

func (e EvolutionDetail) Location(f Finder) (*Location, error) {
	if !e.LocationID.Valid {
		return nil, nil
	}

	var l Location
	query := sq.Select("*").From("locations").Where("id = ?", e.LocationID.Int64)
	if err := f.Find(&l, query); err != nil {
		return nil, err
	}

	return &l, nil
}

func (e EvolutionDetail) HeldItem(f Finder) (*Item, error) {
	if !e.HeldItemID.Valid {
		return nil, nil
	}

	var i Item
	query := sq.Select("*").From("items").Where("id = ?", e.HeldItemID.Int64)
	if err := f.Find(&i, query); err != nil {
		return nil, err
	}

	return &i, nil
}

func (e EvolutionDetail) KnownMove(f Finder) (*Move, error) {
	if !e.KnownMoveID.Valid {
		return nil, nil
	}

	var m Move
	query := sq.Select("*").From("moves").Where("id = ?", e.KnownMoveID.Int64)
	if err := f.Find(&m, query); err != nil {
		return nil, err
	}

	return &m, nil
}

func (e EvolutionDetail) KnownMoveType(f Finder) (*Type, error) {
	if !e.KnownMoveTypeID.Valid {
		return nil, nil
	}

	var t Type
	query := sq.Select("*").From("types").Where("id = ?", e.KnownMoveTypeID.Int64)
	if err := f.Find(&t, query); err != nil {
		return nil, err
	}

	return &t, nil
}

func (e EvolutionDetail) PartySpecies(f Finder) (*PokemonSpecies, error) {
	if !e.PartySpeciesID.Valid {
		return nil, nil
	}

	var s PokemonSpecies
	query := sq.Select("*").From("pokemon_species").Where("id = ?", e.PartySpeciesID.Int64)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (e EvolutionDetail) PartyType(f Finder) (*Type, error) {
	if !e.PartyTypeID.Valid {
		return nil, nil
	}

	var t Type
	query := sq.Select("*").From("types").Where("id = ?", e.PartyTypeID.Int64)
	if err := f.Find(&t, query); err != nil {
		return nil, err
	}

	return &t, nil
}

func (e EvolutionDetail) TradeSpecies(f Finder) (*PokemonSpecies, error) {
	if !e.TradeSpeciesID.Valid {
		return nil, nil
	}

	var s PokemonSpecies
	query := sq.Select("*").From("pokemon_species").Where("id = ?", e.TradeSpeciesID.Int64)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (e EvolutionChain) Details(f Finder) (*EvolutionDetail, error) {
	if !e.EvolutionDetailsID.Valid {
		return nil, nil
	}

	var d EvolutionDetail
	query := sq.Select("pe.*", "t.identifier AS time", "g.identifier AS gender").From("pokemon_evolution as pe").
		LeftJoin("times AS t ON t.id = pe.time_id").
		LeftJoin("genders AS g ON g.id = pe.gender_id").
		Where("id = ?", e.EvolutionDetailsID.Int64)

	if err := f.Find(&d, query); err != nil {
		return nil, err
	}

	return &d, nil
}

func (e EvolutionChain) EvolvesTo(f Finder) (*PokemonSpecies, error) {
	if !e.EvolvesToSpeciesID.Valid {
		return nil, nil
	}

	var s PokemonSpecies
	query := sq.Select("*").From("pokemon_species").Where("id = ?", e.EvolvesToSpeciesID.Int64)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (e EvolutionChain) EvolvesFrom(f Finder) (*PokemonSpecies, error) {
	if !e.EvolvesFromSpeciesID.Valid {
		return nil, nil
	}

	var s PokemonSpecies
	query := sq.Select("*").From("pokemon_species").Where("id = ?", e.EvolvesFromSpeciesID.Int64)
	if err := f.Find(&s, query); err != nil {
		return nil, err
	}

	return &s, nil
}

func (e EvolutionChain) BabyTriggerItem(f Finder) (*Item, error) {
	if !e.BabyTriggerItemID.Valid {
		return nil, nil
	}

	var i Item
	query := sq.Select("*").From("items").Where("id = ?", e.BabyTriggerItemID)
	if err := f.Find(&i, query); err != nil {
		return nil, err
	}

	return &i, nil
}
