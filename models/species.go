package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type PokemonSpecies struct {
	ID                   int64          `db:"id"`
	Identifier           string         `db:"identifier"`
	GenerationID         int64          `db:"generation_id"`
	EvolvesFromSpeciesID sql.NullInt64  `db:"evolves_from_species_id"`
	EvolutionChainID     int64          `db:"evolution_chain_id"`
	ColorID              int64          `db:"color_id"`
	ShapeID              int64          `db:"shape_id"`
	HabitatID            sql.NullInt64  `db:"habitat_id"`
	GenderRate           int64          `db:"gender_rate"`
	CaptureRate          int64          `db:"capture_rate"`
	BaseHappiness        int64          `db:"base_happiness"`
	Baby                 bool           `db:"is_baby"`
	HatchCounter         int64          `db:"hatch_counter"`
	HasGenderDifferences bool           `db:"has_gender_differences"`
	GrowthRateID         int64          `db:"growth_rate_id"`
	FormsSwitchable      bool           `db:"forms_switchable"`
	Ordering             int64          `db:"ordering"`
	FormDescription      sql.NullString `db:"form_description"`
	Name                 string         `db:"name"`
	Genus                string         `db:"genus"`
}

type PokemonShape struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type PokemonHabitat struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

type PokemonColor struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

type PokemonDexNumber struct {
	SpeciesID     int64 `db:"species_id"`
	PokedexID     int64 `db:"pokedex_id"`
	PokedexNumber int64 `db:"pokedex_number"`
}

func (s PokemonSpecies) FlavorText(f Finder, vg int64) (*FlavorText, error) {
	var flav FlavorText
	query := sq.Select("version_group_id", "flavor_text").
		From("pokemon_species_flavor_text").
		Where(sq.Eq{"species_id": s.ID, "version_group_id": vg})

	if err := f.Find(&flav, query); err != nil {
		return nil, err
	}

	return &flav, nil
}

func (s PokemonSpecies) Evolution(f Finder) (*EvolutionChain, error) {
	var c EvolutionChain

	query := EvolutionChains().
		Where(sq.Eq{"evolution_chain_id": s.EvolutionChainID, "ps.id": s.ID})

	if err := f.Find(&c, query); err != nil {
		return nil, err
	}

	return &c, nil
}

func (s PokemonSpecies) Generation(f Finder) (*Generation, error) {
	var gen Generation
	query := sq.Select("*").From("generations").Where("id = ?", s.GenerationID)
	if err := f.Find(&gen, query); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (s PokemonSpecies) EvolvesFrom(f Finder) (*PokemonSpecies, error) {
	if !s.EvolvesFromSpeciesID.Valid {
		return nil, nil
	}

	var ev PokemonSpecies
	query := sq.Select("*").From("pokemon_species").Where("id = ?", s.EvolvesFromSpeciesID.Int64)
	if err := f.Find(&ev, query); err != nil {
		return nil, err
	}

	return &ev, nil
}

func (s PokemonSpecies) Color(f Finder) (*PokemonColor, error) {
	var c PokemonColor
	query := sq.Select("*").From("pokemon_colors").Where("id = ?", s.ColorID)
	if err := f.Find(&c, query); err != nil {
		return nil, err
	}

	return &c, nil
}

func (s PokemonSpecies) Shape(f Finder) (*PokemonShape, error) {
	var sh PokemonShape
	query := sq.Select("*").From("pokemon_shapes").Where("id = ?", s.ShapeID)
	if err := f.Find(&sh, query); err != nil {
		return nil, err
	}

	return &sh, nil
}

func (s PokemonSpecies) Habitat(f Finder) (*PokemonHabitat, error) {
	if !s.HabitatID.Valid {
		return nil, nil
	}

	var h PokemonHabitat
	query := sq.Select("*").From("pokemon_habitats").Where("id = ?", s.HabitatID.Int64)
	if err := f.Find(&h, query); err != nil {
		return nil, err
	}

	return &h, nil
}

func (s PokemonSpecies) Growth(f Finder) (*GrowthRate, error) {
	var r GrowthRate
	query := sq.Select("*").From("growth_rates").Where("id = ?", s.GrowthRateID)
	if err := f.Find(&r, query); err != nil {
		return nil, err
	}

	return &r, nil
}

func (s PokemonSpecies) EggGroups(f Finder) ([]EggGroup, error) {
	var egg []EggGroup
	query := sq.Select("id", "identifier", "name").From("pokemon_egg_groups as peg").
		Join("egg_groups AS eg ON eg.id = peg.egg_group_id").
		Where("species_id = ?", s.ID)

	if err := f.Find(&egg, query); err != nil {
		return nil, err
	}

	return egg, nil
}

func (d PokemonDexNumber) PokeDex(f Finder) (*Pokedex, error) {
	var p Pokedex
	query := sq.Select("*").From("pokedexes").Where("id = ?", d.PokedexID)
	if err := f.Find(&p, query); err != nil {
		return nil, err
	}

	return &p, nil
}
