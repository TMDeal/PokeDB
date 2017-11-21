package regions

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//Resolver resolves the fields of a Region
type Resolver struct {
	db     *models.DB
	region *models.Region
}

//NewResolver returns a new RegionResolver
func NewResolver(db *models.DB, r *models.Region) *Resolver {
	return &Resolver{db, r}
}

//ID resolves the ID field of a Region
func (r *Resolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(r.region.ID)))
	return id
}

//Identifier resolves the Identifier field of a Region
func (r *Resolver) Identifier() string {
	return r.region.Identifier
}

//Name resolves the Name field of a Region
func (r *Resolver) Name() string {
	return r.region.Name
}
