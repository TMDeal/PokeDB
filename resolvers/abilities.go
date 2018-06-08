package resolvers

//go:generate ../connection -model=Ability -table=abilities

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type AbilityResolver struct {
	db *models.DB
	a  *models.Ability
}

func NewAbilityResolver(db *models.DB, a *models.Ability) *AbilityResolver {
	return &AbilityResolver{db, a}
}

func (r AbilityResolver) ID() graphql.ID {
	return GlobalID(models.Ability{}, r.a.ID)
}

func (r AbilityResolver) Identifier() string {
	return r.a.Identifier
}

func (r AbilityResolver) Name() string {
	return r.a.Name
}

func (r AbilityResolver) ShortEffect() string {
	return r.a.ShortEffect
}

func (r AbilityResolver) Effect() string {
	return r.a.Effect
}

func (r AbilityResolver) Generation() (*GenerationResolver, error) {
	gen, err := r.a.Generation(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewGenerationResolver(r.db, gen), nil
}

func (r AbilityResolver) FlavorText(args arguments.FlavorText) (*FlavorTextResolver, error) {
	flav, err := r.a.FlavorText(r.db, int64(args.VersionGroup))
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewFlavorTextResolver(r.db, flav), nil
}
