package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveFlagResolver struct {
	db *models.DB
	mf *models.MoveFlag
}

func NewMoveFlagResolver(db *models.DB, mf *models.MoveFlag) *MoveFlagResolver {
	return &MoveFlagResolver{db, mf}
}

func (r MoveFlagResolver) Identifier() string {
	return r.mf.Identifier
}

func (r MoveFlagResolver) Name() string {
	return r.mf.Name
}

func (r MoveFlagResolver) Description() string {
	return r.mf.Description
}
