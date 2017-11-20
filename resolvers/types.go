package resolvers

import (
	"strconv"

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

func (tr *TypeResolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(tr.t.ID)))
	return id
}

func (tr *TypeResolver) Identifier() string {
	return tr.t.Identifier
}

func (tr *TypeResolver) Name() string {
	return tr.t.Name
}

func (tr *TypeResolver) Generation() *GenerationResolver {
	gen, err := tr.t.Generation(tr.db)
	if err != nil {
		return nil
	}

	return NewGenerationResolver(tr.db, gen)
}

func (tr *TypeResolver) DamageClass() *DamageClassResolver {
	dc, err := tr.t.DamageClass(tr.db)
	if err != nil {
		return nil
	}

	return NewDamageClassResolver(tr.db, dc)
}
