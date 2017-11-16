package db

import "log"

type Version struct {
	ID           int           `db:"id"`
	Name         string        `db:"name"`
	VersionGroup *VersionGroup `db:"version_group"`
}

func (db DB) FindVersionByID(id int) (*Version, error) {
	var v Version

	err := db.QueryRowx(`
	select
	v.id, v.name, v.version_group_id as "version_group.id" ,
	vg.name as "version_group.name", vg.generation_id as "version_group.generation.id",
	g.name as "version_group.generation.name", g.region_id as "version_group.generation.region.id",
	r.name as "version_group.generation.region.name"
	from
	versions as v, version_groups as vg, regions as r, generations as g
	where
	v.version_group_id = vg.id
	and vg.generation_id = g.id
	and g.region_id = r.id
	and v.id = $1
	`, id).StructScan(&v)

	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}

	return &v, nil
}

func (db DB) FindVersionByName(name string) (*Version, error) {
	var v Version

	err := db.QueryRowx(`
	select
	v.id, v.name, v.version_group_id as "version_group.id" ,
	vg.name as "version_group.name", vg.generation_id as "version_group.generation.id",
	g.name as "version_group.generation.name", g.region_id as "version_group.generation.region.id",
	r.name as "version_group.generation.region.name"
	from
	versions as v, version_groups as vg, regions as r, generations as g
	where
	v.version_group_id = vg.id
	and vg.generation_id = g.id
	and g.region_id = r.id
	and v.name = $1
	`, name).StructScan(&v)

	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}

	return &v, nil
}
