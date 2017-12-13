package resolvers

//go:generate go run ./connection/main.go -model=Region -table=regions

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//RegionResolver resolves the fields of a Region
type RegionResolver struct {
	db     *models.DB
	region *models.Region
}

//NewRegionResolver returns a new RegionResolver
func NewRegionResolver(db *models.DB, r *models.Region) *RegionResolver {
	return &RegionResolver{db, r}
}

//ID resolves the ID field of a Region
func (r *RegionResolver) ID() graphql.ID {
	return GlobalID(models.Region{}, r.region.ID)
}

//Identifier resolves the Identifier field of a Region
func (r *RegionResolver) Identifier() string {
	return r.region.Identifier
}

//Name resolves the Name field of a Region
func (r *RegionResolver) Name() string {
	return r.region.Name
}
