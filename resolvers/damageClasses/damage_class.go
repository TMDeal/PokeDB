package damageClasses

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

//Resolver resolves a DamageClass for graphql
type Resolver struct {
	db          *models.DB
	damageClass *models.DamageClass
}

//NewResolver returns a new Resolver
func NewResolver(db *models.DB, dc *models.DamageClass) *Resolver {
	return &Resolver{db, dc}
}

//ID resolves a damage classes ID
func (r *Resolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(r.damageClass.ID)))
	return id
}

//Identifier resolves a damage classes identifier
func (r *Resolver) Identifier() string {
	return r.damageClass.Identifier
}

//Name resolves a damage classes name
func (r *Resolver) Name() string {
	return r.damageClass.Name
}

//Description resolves a damage classes description
func (r *Resolver) Description() string {
	return r.damageClass.Description
}
