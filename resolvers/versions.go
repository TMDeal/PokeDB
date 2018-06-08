package resolvers

//go:generate ../connection -model=Version -table=versions

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type VersionResolver struct {
	db *models.DB
	v  *models.Version
}

func NewVersionResolver(db *models.DB, v *models.Version) *VersionResolver {
	return &VersionResolver{db, v}
}

func (r VersionResolver) ID() graphql.ID {
	return GlobalID(models.Version{}, r.v.ID)
}

func (r VersionResolver) Identifier() string {
	return r.v.Identifier
}

func (r VersionResolver) Name() string {
	return r.v.Name
}

func (r VersionResolver) Group() (*VersionGroupResolver, error) {
	vg, err := r.v.VersionGroup(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewVersionGroupResolver(r.db, vg), nil
}
