package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type TypeResolver struct {
	t *models.Type
}

func NewTypeResolver(t *models.Type) *TypeResolver {
	return &TypeResolver{t}
}

func (tr *TypeResolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(tr.t.ID))
	return id
}

func (tr *TypeResolver) Identifier() string {
	return tr.t.Identifier
}

func (tr *TypeResolver) Name() string {
	return tr.t.Name
}

func (tr *TypeResolver) Generation() *GenerationResolver {
	gen, err := tr.t.Generation()
	if err != nil {
		return nil
	}

	return NewGenerationResolver(gen)
}

func (tr *TypeResolver) DamageClass() *DamageClassResolver {
	dc, err := tr.t.DamageClass()
	if err != nil {
		return nil
	}

	return NewDamageClassResolver(dc)
}
