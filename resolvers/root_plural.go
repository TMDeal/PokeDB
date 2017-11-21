package resolvers

import (
	"database/sql"
	"log"

	"github.com/TMDeal/PokeDB/resolvers/generations"
	"github.com/TMDeal/PokeDB/resolvers/regions"
	"github.com/TMDeal/PokeDB/resolvers/types"
)

//Generation resolves a Generation based on an ID
func (root *RootResolver) Generations(args struct{ First int32 }) *[]*generations.Resolver {
	var gr []*generations.Resolver

	gens, err := root.db.FindGenerations(uint64(args.First))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, gen := range gens {
		gr = append(gr, generations.NewResolver(root.db, gen))
	}

	return &gr
}

//Region resolves a Region based on an ID
func (root *RootResolver) Regions(args struct{ First int32 }) *[]*regions.Resolver {
	var rr []*regions.Resolver

	rs, err := root.db.FindRegions(uint64(args.First))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rs {
		rr = append(rr, regions.NewResolver(root.db, r))
	}

	return &rr
}

//Types resolves a Type based on an ID
func (root *RootResolver) Types(args struct{ First int32 }) *[]*types.Resolver {
	var tr []*types.Resolver

	ts, err := root.db.FindTypes(uint64(args.First))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range ts {
		tr = append(tr, types.NewResolver(root.db, t))
	}

	return &tr
}
