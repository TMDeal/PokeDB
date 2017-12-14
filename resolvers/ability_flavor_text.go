package resolvers

import "github.com/TMDeal/PokeDB/models"

type AbilityFlavorTextResolver struct {
	db   *models.DB
	flav *models.AbilityFlavorText
}

func NewAbilityFlavorTextResolver(db *models.DB, flav *models.AbilityFlavorText) *AbilityFlavorTextResolver {
	return &AbilityFlavorTextResolver{db, flav}
}

func (r AbilityFlavorTextResolver) Text() string {
	return r.flav.Text
}

func (r AbilityFlavorTextResolver) VersionGroup() (*VersionGroupResolver, error) {
	vg, err := r.flav.VersionGroup(r.db)
	if err != nil {
		return nil, err
	}

	return NewVersionGroupResolver(r.db, vg), nil
}
