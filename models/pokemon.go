package models

import "log"

type Pokemon struct {
	ID           int                    `db:"id"`
	Identifier   string                 `db:"identifier"`
	Height       float32                `db:"height"`
	Weight       float32                `db:"weight"`
	MaleChance   float32                `db:"male_chance"`
	FemaleChance float32                `db:"female_chance"`
	Color        string                 `db:"color"`
	Types        []PokemonType          `db:"types"`
	Stats        map[string]PokemonStat `db:"stats"`
}

type PokemonStat struct {
	Stat
	Effort int `db:"effort"`
	Base   int `db:"base"`
}

type PokemonType struct {
	Type
}

func (db DB) FindPokemonByID(id int) (*Pokemon, error) {
	var pokemon Pokemon
	var types []PokemonType
	stats := make(map[string]PokemonStat)

	err := db.QueryRowx(`
	select * from pokemon where id = $1
	`, id).StructScan(&pokemon)
	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}

	rowsStat, err := db.Queryx(`
	select s.identifier, s.alt_identifier, s.id, ps.effort, ps.base
	from pokemon_stats as ps, stats as s
	where ps.stat_id = s.id and ps.pokemon_id = $1
	`, id)
	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}
	defer rowsStat.Close()

	for rowsStat.Next() {
		var stat PokemonStat
		rowsStat.StructScan(&stat)
		stats[stat.AltIdentifier] = stat
	}

	rowsType, err := db.Queryx(`
	select t.id, t.identifier
	from types as t, pokemon_types as pt
	where pt.type_id = t.id and pt.pokemon_id = $1
	`, id)
	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}
	defer rowsType.Close()

	for rowsType.Next() {
		var t PokemonType
		rowsType.StructScan(&t)
		types = append(types, t)
	}

	pokemon.Types = types
	pokemon.Stats = stats

	return &pokemon, nil
}
