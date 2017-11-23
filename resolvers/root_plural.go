package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/gocraft/dbr"
)

func (root *RootResolver) Generations(args arguments.Connection) GenerationConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	items, err := root.db.FindGenerations(limit, offset)
	if err != nil && err != dbr.ErrNotFound {
		log.Fatal(err)
	}

	connections, err := NewGenerationConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Regions(args arguments.Connection) RegionConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	items, err := root.db.FindRegions(limit, offset)
	if err != nil && err != dbr.ErrNotFound {
		log.Fatal(err)
	}

	connections, err := NewRegionConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Types(args arguments.Connection) TypeConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	items, err := root.db.FindTypes(limit, offset)
	if err != nil && err != dbr.ErrNotFound {
		log.Fatal(err)
	}

	connections, err := NewTypeConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}
