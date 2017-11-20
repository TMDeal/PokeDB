package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//RegionResolver resolves the fields of a Region
type RegionResolver struct {
	db *models.DB
	r  *models.Region
}

//NewRegionResolver returns a new RegionResolver
func NewRegionResolver(db *models.DB, r *models.Region) *RegionResolver {
	return &RegionResolver{db, r}
}

//ID resolves the ID field of a Region
func (rr *RegionResolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(rr.r.ID)))
	return id
}

//Identifier resolves the Identifier field of a Region
func (rr *RegionResolver) Identifier() string {
	return rr.r.Identifier
}

//Name resolves the Name field of a Region
func (rr *RegionResolver) Name() string {
	return rr.r.Name
}
