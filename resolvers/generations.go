package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//GenerationResolver resolves the fields of a Generation
type GenerationResolver struct {
	db *models.DB
	g  *models.Generation
}

//NewGenerationResolver returns a new GenerationResolver
func NewGenerationResolver(db *models.DB, g *models.Generation) *GenerationResolver {
	return &GenerationResolver{db, g}
}

//ID resolves the ID field of a Generation
func (gr *GenerationResolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(gr.g.ID))
	return id
}

//Identifier resolves the Identifier field of a Generation
func (gr *GenerationResolver) Identifier() string {
	return gr.g.Identifier
}

//Name resolves the Name field of a Generation
func (gr *GenerationResolver) Name() string {
	return gr.g.Name
}

//Region resolves the Region of a generation, by finding the region based on a
//Generations RegionID, and returning a RegionResolver for that Region
func (gr *GenerationResolver) Region() *RegionResolver {
	r, err := gr.g.Region(gr.db)
	if err != nil {
		return nil
	}
	return NewRegionResolver(gr.db, r)
}
