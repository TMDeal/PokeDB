package resolvers

import "github.com/TMDeal/PokeDB/models"

//RootResolver is the root resolver for the graphql schema. All other resolvers
//get returned from this resolver
type RootResolver struct {
	db *models.DB
}

//NewRootResolver returns a new RootResolver
func NewRootResolver(db *models.DB) *RootResolver {
	return &RootResolver{db}
}
