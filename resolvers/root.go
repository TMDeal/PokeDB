package resolvers

import (
	"database/sql"
	"log"

	"github.com/TMDeal/PokeDB/models"
)

//RootResolver is the root resolver for the graphql schema. All other resolvers
//get returned from this resolver
type RootResolver struct {
	DB *models.DB
}

//NewRootResolver returns a new RootResolver
func NewRootResolver(db *models.DB) *RootResolver {
	return &RootResolver{
		DB: db,
	}
}

type IDArgs struct {
	ID *int32
}

type IDNameArgs struct {
	ID   *int32
	Name *string
}

//Generation resolves a Generation based on an ID
func (root *RootResolver) Generations(args IDArgs) *[]*GenerationResolver {
	var gr []*GenerationResolver

	gens, err := root.DB.FindGenerations(int(*args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, gen := range gens {
		gr = append(gr, NewGenerationResolver(gen))
	}

	return &gr
}

//Region resolves a Region based on an ID
func (root *RootResolver) Regions(args IDNameArgs) *[]*RegionResolver {
	var rr []*RegionResolver

	rs, err := root.DB.FindRegions(int(*args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rs {
		rr = append(rr, NewRegionResolver(r))
	}

	return &rr
}

//Types resolves a Type based on an ID
func (root *RootResolver) Types(args IDNameArgs) *[]*TypeResolver {
	var tr []*TypeResolver

	ts, err := root.DB.FindTypes(int(*args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range ts {
		tr = append(tr, NewTypeResolver(t))
	}
	return &tr
}
