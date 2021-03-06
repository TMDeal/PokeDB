package resolvers

//go:generate ../connection -model=VersionGroup -table=version_groups

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type VersionGroupResolver struct {
	db *models.DB
	vg *models.VersionGroup
}

func NewVersionGroupResolver(db *models.DB, vg *models.VersionGroup) *VersionGroupResolver {
	return &VersionGroupResolver{db, vg}
}

func (r VersionGroupResolver) ID() graphql.ID {
	return GlobalID(models.VersionGroup{}, r.vg.ID)
}

func (r VersionGroupResolver) Identifier() string {
	return r.vg.Identifier
}

func (r VersionGroupResolver) Ordering() int32 {
	return int32(r.vg.Ordering)
}

func (r VersionGroupResolver) Generation() (*GenerationResolver, error) {
	gen, err := r.vg.Generation(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewGenerationResolver(r.db, gen), nil
}

func (r VersionGroupResolver) Versions() ([]*VersionResolver, error) {
	vs, err := r.vg.Versions(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	var vrs []*VersionResolver
	for i, _ := range vs {
		vrs = append(vrs, NewVersionResolver(r.db, &vs[i]))
	}

	return vrs, nil
}

func (r VersionGroupResolver) Regions() ([]*RegionResolver, error) {
	rs, err := r.vg.Regions(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	var rrs []*RegionResolver
	for i, _ := range rs {
		rrs = append(rrs, NewRegionResolver(r.db, &rs[i]))
	}

	return rrs, nil
}
