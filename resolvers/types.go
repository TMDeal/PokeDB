package resolvers

//go:generate go run ./connection/main.go -model=Type -table=types

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type TypeResolver struct {
	db *models.DB
	t  *models.Type
}

func NewTypeResolver(db *models.DB, t *models.Type) *TypeResolver {
	return &TypeResolver{db, t}
}

func (r *TypeResolver) ID() graphql.ID {
	return GlobalID(models.Type{}, r.t.ID)
}

func (r *TypeResolver) Identifier() string {
	return r.t.Identifier
}

func (r *TypeResolver) Name() string {
	return r.t.Name
}

func (r *TypeResolver) Generation() *GenerationResolver {
	gen, err := r.t.Generation(r.db)
	if err != nil {
		return nil
	}

	return NewGenerationResolver(r.db, gen)
}

func (r *TypeResolver) DamageClass() *DamageClassResolver {
	dc, err := r.t.DamageClass(r.db)
	if err != nil {
		return nil
	}

	return NewDamageClassResolver(r.db, dc)
}

func (r *TypeResolver) Efficacy() *TypeEfficacyResolver {
	return NewTypeEfficacyResolver(r.db, r.t)
}
