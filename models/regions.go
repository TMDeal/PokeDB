package models

import (
	"github.com/gocraft/dbr"
)

//Region represents a region entry in the database
type Region struct {
	ID         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Name       string `db:"name"`
}

//RegionFinder says how to find information for a region model
type RegionFinder interface {
	FindRegions(limit uint64) ([]*Region, error)
	FindRegion(query string, value interface{}) (*Region, error)
}

func (db DB) FindRegion(query string, value interface{}) (*Region, error) {
	var r Region
	sess := db.Session()

	count, err := sess.Select("*").From("regions").Where(query, value).Load(&r)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &r, nil
}

func (db DB) FindRegions(limit uint64) ([]*Region, error) {
	var rs []*Region
	sess := db.Session()

	count, err := sess.Select("*").From("regions").Limit(limit).Load(&rs)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return rs, nil
}
