package db

import "log"

type VersionGroup struct {
	ID         int         `db:"id"`
	Name       string      `db:"name"`
	Generation *Generation `db:"generation"`
}

func (db DB) FindVersionGroupByID(id int) (*VersionGroup, error) {
	var vg VersionGroup

	err := db.QueryRowx(`
	select
	vg.id, vg.name, vg.generation_id as "generation.id",
	g.name as "generation.name", g.region_id as "generation.region.id",
	r.name as "generation.region.name"
	from version_groups as vg, generations as g, regions as r
	where
	vg.generation_id = g.id
	and g.region_id = r.id
	and vg.id = $1
	`, id).StructScan(&vg)

	if err != nil {
		log.Fatal("Unable to execute query!")
		return nil, err
	}

	return &vg, nil
}

func (db DB) FindVersionGroupByName(name string) (*VersionGroup, error) {
	var vg VersionGroup

	err := db.QueryRowx(`
	select
	vg.id, vg.name, vg.generation_id as "generation.id",
	g.name as "generation.name", g.region_id as "generation.region.id",
	r.name as "generation.region.name"
	from version_groups as vg, generations as g, regions as r
	where
	vg.generation_id = g.id
	and g.region_id = r.id
	and vg.name = $1
	`, name).StructScan(&vg)

	if err != nil {
		log.Fatal("Unable to execute query!")
		return nil, err
	}

	return &vg, nil
}