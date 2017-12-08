package resolvers

import "github.com/TMDeal/PokeDB/models"

type ContestTypeResolver struct {
	db *models.DB
	ct *models.ContestType
}

func NewContestTypeResolver(db *models.DB, ct *models.ContestType) *ContestTypeResolver {
	return &ContestTypeResolver{db, ct}
}

func (r ContestTypeResolver) Identifier() string {
	return r.ct.Identifier
}

func (r ContestTypeResolver) Name() string {
	return r.ct.Name
}

func (r ContestTypeResolver) Flavor() string {
	return r.ct.Flavor
}

func (r ContestTypeResolver) Color() string {
	return r.ct.Color
}
