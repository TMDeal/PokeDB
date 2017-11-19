package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/models"
)

//RootResolver is the root resolver for the graphql schema. All other resolvers
//get returned from this resolver
type RootResolver struct {
	DB *models.DB
}

//NewRootResolver returns a new RootResolver
func NewRootResolver(db *models.DB) *RootResolver {
	return &RootResolver{
		DB: db,
	}
}

//Generation resolves a Generation based on an ID
func (root *RootResolver) Generation(args struct{ ID int32 }) *GenerationResolver {
	gen, err := root.DB.FindGenerationByID(int(args.ID))
	if err != nil {
		log.Fatal(err)
	}
	return NewGenerationResolver(gen)
}

//Region resolves a Region based on an ID
func (root *RootResolver) Region(args struct{ ID int32 }) *RegionResolver {
	r, err := root.DB.FindRegionByID(int(args.ID))
	if err != nil {
		log.Fatal(err)
	}
	return NewRegionResolver(r)
}
