package resolvers

//go:generate ../connection -model=Berry -table=berries

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type BerryResolver struct {
	*ItemResolver
	db *models.DB
	b  *models.Berry
}

func NewBerryResolver(db *models.DB, b *models.Berry) *BerryResolver {
	return &BerryResolver{NewItemResolver(db, &b.Item), db, b}
}

func (r BerryResolver) ID() graphql.ID {
	return GlobalID(models.Berry{}, r.b.ID)
}

func (r BerryResolver) ItemID() graphql.ID {
	return r.ItemResolver.ID()
}

func (r BerryResolver) NaturalGiftPower() int32 {
	return int32(r.b.NaturalGiftPower)
}

func (r BerryResolver) Size() int32 {
	return int32(r.b.Size)
}

func (r BerryResolver) MaxHarvest() int32 {
	return int32(r.b.MaxHarvest)
}

func (r BerryResolver) GrowthTime() int32 {
	return int32(r.b.GrowthTime)
}

func (r BerryResolver) SoilDryness() int32 {
	return int32(r.b.SoilDryness)
}

func (r BerryResolver) Smoothness() int32 {
	return int32(r.b.Smoothness)
}

func (r BerryResolver) Firmness() (*BerryFirmnessResolver, error) {
	firm, err := r.b.Firmness(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewBerryFirmnessResolver(r.db, firm), nil
}

func (r BerryResolver) Flavors() ([]*BerryFlavorResolver, error) {
	flavs, err := r.b.Flavors(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	var flavsr []*BerryFlavorResolver
	for i, _ := range flavs {
		flavsr = append(flavsr, NewBerryFlavorResolver(r.db, &flavs[i]))
	}

	return flavsr, nil
}
