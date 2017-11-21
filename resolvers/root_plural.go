package resolvers

import (
	"database/sql"
	"log"
)

//Generation resolves a Generation based on an ID
func (root *RootResolver) Generations(args struct{ First int32 }) *[]*GenerationResolver {
	var gr []*GenerationResolver

	gens, err := root.db.FindGenerations(uint64(args.First))
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
func (root *RootResolver) Regions(args struct{ First int32 }) *[]*RegionResolver {
	var rr []*RegionResolver

	rs, err := root.db.FindRegions(uint64(args.First))
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
func (root *RootResolver) Types(args struct{ First int32 }) *[]*TypeResolver {
	var tr []*TypeResolver

	ts, err := root.db.FindTypes(uint64(args.First))
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
