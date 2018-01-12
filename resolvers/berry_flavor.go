package resolvers

import "github.com/TMDeal/PokeDB/models"

type BerryFlavorResolver struct {
	db *models.DB
	bf *models.BerryFlavor
}

func NewBerryFlavorResolver(db *models.DB, bf *models.BerryFlavor) *BerryFlavorResolver {
	return &BerryFlavorResolver{db, bf}
}

func (r BerryFlavorResolver) Name() string {
	return r.bf.Name
}

func (r BerryFlavorResolver) Potency() int32 {
	return int32(r.bf.Potency)
}

func (r BerryFlavorResolver) ContestType() (*ContestTypeResolver, error) {
	ct, err := r.bf.ContestType(r.db)
	if err != nil {
		return nil, err
	}

	return NewContestTypeResolver(r.db, ct), nil
}
