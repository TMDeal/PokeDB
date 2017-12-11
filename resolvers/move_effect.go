package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveEffectResolver struct {
	db *models.DB
	me *models.MoveEffect
}

func NewMoveEffectResolver(db *models.DB, me *models.MoveEffect) *MoveEffectResolver {
	return &MoveEffectResolver{db, me}
}

func (r *MoveEffectResolver) Short() string {
	return r.me.ShortEffect
}

func (r *MoveEffectResolver) Long() string {
	return r.me.Effect
}
