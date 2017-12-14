package resolvers

//go:generate go run ./connection/main.go -model=Ability -table=abilities

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
		return nil, err
	}

	return NewGenerationResolver(r.db, gen), nil
}

func (r AbilityResolver) FlavorText(args arguments.FlavorText) (string, error) {
	mft, err := r.a.FlavorText(r.db, int64(args.VersionGroup))
	if err != nil {
		return "", err
	}

	return mft.Text, nil
}
