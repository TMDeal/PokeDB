package resolvers

import "github.com/TMDeal/PokeDB/models"

type FlavorTextResolver struct {
	db   *models.DB
	flav *models.FlavorText
}

func NewFlavorTextResolver(db *models.DB, flav *models.FlavorText) *FlavorTextResolver {
	return &FlavorTextResolver{db, flav}
}

func (r FlavorTextResolver) Text() string {
	return r.flav.Text
}

func (r FlavorTextResolver) VersionGroup() (*VersionGroupResolver, error) {
	vg, err := r.flav.VersionGroup(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewVersionGroupResolver(r.db, vg), nil
}
