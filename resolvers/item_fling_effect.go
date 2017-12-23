package resolvers

import "github.com/TMDeal/PokeDB/models"

type ItemFlingEffectResolver struct {
	db  *models.DB
	ife *models.ItemFlingEffect
}

func NewItemFlingEffectResolver(db *models.DB, ife *models.ItemFlingEffect) *ItemFlingEffectResolver {
	return &ItemFlingEffectResolver{db, ife}
}

func (r ItemFlingEffectResolver) Identifier() string {
	return r.ife.Identifier
}

func (r ItemFlingEffectResolver) Effect() string {
	return r.ife.Effect
}
