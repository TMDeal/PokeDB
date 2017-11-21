package generations

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/resolvers/regions"
	graphql "github.com/neelance/graphql-go"
)

//Resolver resolves the fields of a Generation
type Resolver struct {
	db         *models.DB
	generation *models.Generation
}

//NewResolver returns a new Resolver
func NewResolver(db *models.DB, g *models.Generation) *Resolver {
	return &Resolver{db, g}
}

//ID resolves the ID field of a Generation
func (r *Resolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(r.generation.ID)))
	return id
}

//Identifier resolves the Identifier field of a Generation
func (r *Resolver) Identifier() string {
	return r.generation.Identifier
}

//Name resolves the Name field of a Generation
func (r *Resolver) Name() string {
	return r.generation.Name
}

//Region resolves the Region of a generation, by finding the region based on a
//Generations RegionID, and returning a RegionResolver for that Region
func (r *Resolver) Region() *regions.Resolver {
	rn, err := r.generation.Region(r.db)
	if err != nil {
		return nil
	}
	return regions.NewResolver(r.db, rn)
}
