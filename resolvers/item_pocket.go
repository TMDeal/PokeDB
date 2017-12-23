package resolvers

import "github.com/TMDeal/PokeDB/models"

type ItemPocketResolver struct {
	db *models.DB
	ip *models.ItemPocket
}

func NewItemPocketResolver(db *models.DB, ip *models.ItemPocket) *ItemPocketResolver {
	return &ItemPocketResolver{db, ip}
}

func (r ItemPocketResolver) Identifier() string {
	return r.ip.Identifier
}

func (r ItemPocketResolver) Name() string {
	return r.ip.Name
}
