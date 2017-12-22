package resolvers

//go:generate ../connection -model=Generation -table=generations

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//GenerationResolver resolves the fields of a Generation
type GenerationResolver struct {
	db         *models.DB
	generation *models.Generation
}

//NewGenerationResolver returns a new Resolver
func NewGenerationResolver(db *models.DB, g *models.Generation) *GenerationResolver {
	return &GenerationResolver{db, g}
}

//ID resolves the ID field of a Generation
func (r *GenerationResolver) ID() graphql.ID {
	return GlobalID(models.Generation{}, r.generation.ID)
}

//Identifier resolves the Identifier field of a Generation
func (r *GenerationResolver) Identifier() string {
	return r.generation.Identifier
}

//Name resolves the Name field of a Generation
func (r *GenerationResolver) Name() string {
	return r.generation.Name
}

//Region resolves the Region of a generation, by finding the region based on a
//Generations RegionID, and returning a RegionResolver for that Region
func (r *GenerationResolver) Region() *RegionResolver {
	rn, err := r.generation.Region(r.db)
	if err != nil {
		return nil
	}
	return NewRegionResolver(r.db, rn)
}
