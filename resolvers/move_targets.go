package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveTargetResolver struct {
	db *models.DB
	mf *models.MoveTarget
}

func NewMoveTargetResolver(db *models.DB, mf *models.MoveTarget) *MoveTargetResolver {
	return &MoveTargetResolver{db, mf}
}

func (r MoveTargetResolver) Identifier() string {
	return r.mf.Identifier
}

func (r MoveTargetResolver) Name() string {
	return r.mf.Name
}

func (r MoveTargetResolver) Description() string {
	return r.mf.Description
}
