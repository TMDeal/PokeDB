package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//DamageClassResolver resolves a DamageClass for graphql
type DamageClassResolver struct {
	db *models.DB
	dc *models.DamageClass
}

func NewDamageClassResolver(db *models.DB, dc *models.DamageClass) *DamageClassResolver {
	return &DamageClassResolver{db, dc}
}

//ID resolves a damage classes ID
func (dcr *DamageClassResolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(dcr.dc.ID))
	return id
}

//Identifier resolves a damage classes identifier
func (dcr *DamageClassResolver) Identifier() string {
	return dcr.dc.Identifier
}

//Name resolves a damage classes name
func (dcr *DamageClassResolver) Name() string {
	return dcr.dc.Name
}

//Description resolves a damage classes description
func (dcr *DamageClassResolver) Description() string {
	return dcr.dc.Description
}
