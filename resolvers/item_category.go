package resolvers

import "github.com/TMDeal/PokeDB/models"

type ItemCategoryResolver struct {
	db *models.DB
	ic *models.ItemCategory
}

func NewItemCategoryResolver(db *models.DB, ic *models.ItemCategory) *ItemCategoryResolver {
	return &ItemCategoryResolver{db, ic}
}

func (r ItemCategoryResolver) Identifier() string {
	return r.ic.Identifier
}

func (r ItemCategoryResolver) Name() string {
	return r.ic.Name
}

func (r ItemCategoryResolver) Pocket() (*ItemPocketResolver, error) {
	ip, err := r.ic.Pocket(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewItemPocketResolver(r.db, ip), nil
}
