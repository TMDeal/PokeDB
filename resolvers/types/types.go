package types

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/resolvers/damageClasses"
	"github.com/TMDeal/PokeDB/resolvers/generations"
	graphql "github.com/neelance/graphql-go"
)

type Resolver struct {
	db *models.DB
	t  *models.Type
}

func NewResolver(db *models.DB, t *models.Type) *Resolver {
	return &Resolver{db, t}
}

func (r *Resolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(r.t.ID)))
	return id
}

func (r *Resolver) Identifier() string {
	return r.t.Identifier
}

func (r *Resolver) Name() string {
	return r.t.Name
}

func (r *Resolver) Generation() *generations.Resolver {
	gen, err := r.t.Generation(r.db)
	if err != nil {
		return nil
	}

	return generations.NewResolver(r.db, gen)
}

func (r *Resolver) DamageClass() *damageClasses.Resolver {
	dc, err := r.t.DamageClass(r.db)
	if err != nil {
		return nil
	}

	return damageClasses.NewResolver(r.db, dc)
}
