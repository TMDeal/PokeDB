package resolvers

import "github.com/TMDeal/PokeDB/models"

type BerryFirmnessResolver struct {
	db *models.DB
	bf *models.BerryFirmness
}

func NewBerryFirmnessResolver(db *models.DB, bf *models.BerryFirmness) *BerryFirmnessResolver {
	return &BerryFirmnessResolver{db, bf}
}

func (r BerryFirmnessResolver) Identifier() string {
	return r.bf.Identifier
}

func (r BerryFirmnessResolver) Name() string {
	return r.bf.Name
}
