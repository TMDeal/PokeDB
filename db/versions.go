package db

import "log"

type Version struct {
	ID           int           `db:"id"`
	Identifier   string        `db:"identifier"`
	VersionGroup *VersionGroup `db:"version_group"`
}

func (db DB) FindVersionByID(id int) (*Version, error) {
	var v Version

	err := db.QueryRowx(`
	select
	v.id, v.identifier, v.version_group_id as "version_group.id" ,
	vg.identifier as "version_group.identifier", vg.generation_id as "version_group.generation.id",
	g.identifier as "version_group.generation.identifier", g.region_id as "version_group.generation.region.id",
	r.identifier as "version_group.generation.region.identifier"
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
