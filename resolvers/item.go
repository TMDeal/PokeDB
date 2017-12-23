package resolvers

//go:generate ../connection -model=Item -table=items

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type ItemResolver struct {
	db *models.DB
	i  *models.Item
}

func NewItemResolver(db *models.DB, i *models.Item) *ItemResolver {
	return &ItemResolver{db, i}
}

func (r ItemResolver) ID() graphql.ID {
	return GlobalID(models.Item{}, r.i.ID)
}

func (r ItemResolver) Identifier() string {
	return r.i.Identifier
}

func (r ItemResolver) Name() string {
	return r.i.Name
}

func (r ItemResolver) ShortEffect() string {
	return r.i.ShortEffect
}

func (r ItemResolver) Effect() string {
	return r.i.Effect
}

func (r ItemResolver) Cost() int32 {
	return int32(r.i.Cost)
}

func (r ItemResolver) FlingPower() *int32 {
	if !r.i.FlingPower.Valid {
		return nil
	}

	fp := int32(r.i.FlingPower.Int64)
	return &fp
}

func (r ItemResolver) FlingEffect() (*ItemFlingEffectResolver, error) {
	if !r.i.FlingEffectID.Valid {
		return nil, nil
	}

	ife, err := r.i.FlingEffect(r.db)
	if err != nil {
		return nil, err
	}

	return NewItemFlingEffectResolver(r.db, ife), nil
}

func (r ItemResolver) Category() (*ItemCategoryResolver, error) {
	cat, err := r.i.Category(r.db)
	if err != nil {
		return nil, err
	}

	return NewItemCategoryResolver(r.db, cat), nil
}

func (r ItemResolver) Flags() ([]*ItemFlagResolver, error) {
	fs, err := r.i.Flags(r.db)
	if err != nil {
		return nil, err
	}

	var fsr []*ItemFlagResolver
	for _, f := range fs {
		fsr = append(fsr, NewItemFlagResolver(r.db, &f))
	}

	return fsr, nil
}

func (r ItemResolver) FlavorText(args arguments.FlavorText) (*FlavorTextResolver, error) {
	flav, err := r.i.FlavorText(r.db, int(args.VersionGroup))
	if err != nil {
		return nil, err
	}

	return NewFlavorTextResolver(r.db, flav), nil
}
