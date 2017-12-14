package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveFlavorTextResolver struct {
	db   *models.DB
	flav *models.MoveFlavorText
}

func NewMoveFlavorTextResolver(db *models.DB, flav *models.MoveFlavorText) *MoveFlavorTextResolver {
	return &MoveFlavorTextResolver{db, flav}
}

func (r MoveFlavorTextResolver) Text() string {
	return r.flav.Text
}

func (r MoveFlavorTextResolver) VersionGroup() (*VersionGroupResolver, error) {
	vg, err := r.flav.VersionGroup(r.db)
	if err != nil {
		return nil, err
	}

	return NewVersionGroupResolver(r.db, vg), nil
}
