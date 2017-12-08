package resolvers

import "github.com/TMDeal/PokeDB/models"

type SuperContestEffectResolver struct {
	db  *models.DB
	sce *models.SuperContestEffect
}

func NewSuperContestEffectResolver(db *models.DB, sce *models.SuperContestEffect) *SuperContestEffectResolver {
	return &SuperContestEffectResolver{db, sce}
}

func (r SuperContestEffectResolver) FlavorText() string {
	return r.sce.FlavorText
}

func (r SuperContestEffectResolver) Appeal() string {
	return r.sce.Appeal
}
