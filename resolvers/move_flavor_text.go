package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveFlavorTextResolver struct {
	db *models.DB
	mf *models.MoveFlavorText
}

func NewMoveFlavorTextResolver(db *models.DB, mf *models.MoveFlavorText) *MoveFlavorTextResolver {
	return &MoveFlavorTextResolver{db, mf}
}

