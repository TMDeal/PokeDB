package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveMetaCategoryResolver struct {
	db  *models.DB
	cat *models.MoveMetaCategory
}

func NewMoveMetaCategoryResolver(db *models.DB, cat *models.MoveMetaCategory) *MoveMetaCategoryResolver {
	return &MoveMetaCategoryResolver{db, cat}
}

func (r MoveMetaCategoryResolver) Identifier() string {
	return r.cat.Identifier
}

func (r MoveMetaCategoryResolver) Description() string {
	return r.cat.Description
}
