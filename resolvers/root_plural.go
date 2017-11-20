package resolvers

import (
	"database/sql"
	"log"

	"github.com/TMDeal/PokeDB/models"
)

//RootResolver is the root resolver for the graphql schema. All other resolvers
//get returned from this resolver
type RootResolver struct {
	db *models.DB
}

//NewRootResolver returns a new RootResolver
func NewRootResolver(db *models.DB) *RootResolver {
	return &RootResolver{db}
}

//Generation resolves a Generation based on an ID
func (root *RootResolver) Generations(args struct{ ID *int32 }) *[]*GenerationResolver {
	var gr []*GenerationResolver

	gens, err := root.db.FindGenerations(int(*args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, gen := range gens {
		gr = append(gr, NewGenerationResolver(root.db, gen))
	}

	return &gr
}

//Region resolves a Region based on an ID
func (root *RootResolver) Regions(args struct {
	ID   *int32
	Name *string
}) *[]*RegionResolver {
	var rr []*RegionResolver

	rs, err := root.db.FindRegions(int(*args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rs {
		rr = append(rr, NewRegionResolver(root.db, r))
	}

	return &rr
}

//Types resolves a Type based on an ID
func (root *RootResolver) Types(args struct {
	ID   *int32
	Name *string
}) *[]*TypeResolver {
	var tr []*TypeResolver

	ts, err := root.db.FindTypes(int(*args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range ts {
		tr = append(tr, NewTypeResolver(root.db, t))
	}
	return &tr
}
