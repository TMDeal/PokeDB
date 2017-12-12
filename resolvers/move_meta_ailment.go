package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveMetaAilmentResolver struct {
	db  *models.DB
	ail *models.MoveMetaAilment
}

func NewMoveMetaAilmentResolver(db *models.DB, ail *models.MoveMetaAilment) *MoveMetaAilmentResolver {
	return &MoveMetaAilmentResolver{db, ail}
}

func (r MoveMetaAilmentResolver) Identifier() string {
	return r.ail.Identifier
}

func (r MoveMetaAilmentResolver) Name() string {
	return r.ail.Name
}
