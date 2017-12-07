package resolvers

import (
	"github.com/TMDeal/PokeDB/models"
)

//DamageClassResolver resolves a DamageClass for graphql
type DamageClassResolver struct {
	db          *models.DB
	damageClass *models.DamageClass
}

//NewDamageClassResolver returns a new DamageClassResolver
func NewDamageClassResolver(db *models.DB, dc *models.DamageClass) *DamageClassResolver {
	return &DamageClassResolver{db, dc}
}

//Identifier resolves a damage classes identifier
func (r *DamageClassResolver) Identifier() string {
	return r.damageClass.Identifier
}

//Name resolves a damage classes name
func (r *DamageClassResolver) Name() string {
	return r.damageClass.Name
}

//Description resolves a damage classes description
func (r *DamageClassResolver) Description() string {
	return r.damageClass.Description
}
