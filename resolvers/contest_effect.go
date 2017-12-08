package resolvers

import "github.com/TMDeal/PokeDB/models"

type ContestEffectResolver struct {
	db *models.DB
	ce *models.ContestEffect
}

func NewContestEffectResolver(db *models.DB, ce *models.ContestEffect) *ContestEffectResolver {
	return &ContestEffectResolver{db, ce}
}

func (r ContestEffectResolver) Jam() int32 {
	return int32(r.ce.Jam)
}

func (r ContestEffectResolver) Appeal() int32 {
	return int32(r.ce.Appeal)
}

func (r ContestEffectResolver) FlavorText() string {
	return r.ce.FlavorText
}

func (r ContestEffectResolver) Effect() string {
	return r.ce.Effect
}
