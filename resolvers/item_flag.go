package resolvers

import "github.com/TMDeal/PokeDB/models"

type ItemFlagResolver struct {
	db   *models.DB
	flag *models.ItemFlag
}

func NewItemFlagResolver(db *models.DB, flag *models.ItemFlag) *ItemFlagResolver {
	return &ItemFlagResolver{db, flag}
}

func (r ItemFlagResolver) Identifier() string {
	return r.flag.Identifier
}

func (r ItemFlagResolver) Name() string {
	return r.flag.Name
}

func (r ItemFlagResolver) Description() string {
	return r.flag.Description
}
