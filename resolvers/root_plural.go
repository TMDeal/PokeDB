package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/resolvers/generations"
	"github.com/TMDeal/PokeDB/resolvers/regions"
	"github.com/TMDeal/PokeDB/resolvers/types"
	"github.com/gocraft/dbr"
)

func (root *RootResolver) Generations(args arguments.Connection) generations.ConnectionResolver {
	offset := 0
	limit := 20
	var err error

	if args.After != nil {
		offset, err = args.After.IntValue()
		if err != nil {
			log.Fatal(err)
		}
	}

	if args.First != nil {
		limit = int(*args.First)
	}

	rs, err := root.db.FindGenerations(uint64(limit), uint64(offset))
	if err != nil && err != dbr.ErrNotFound {
		log.Fatal(err)
	}

	connections, err := generations.NewConnectionResolver(root.db, rs, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Regions(args arguments.Connection) regions.ConnectionResolver {
	offset := 0
	limit := 20
	var err error

	if args.After != nil {
		offset, err = args.After.IntValue()
		if err != nil {
			log.Fatal(err)
		}
	}

	if args.First != nil {
		limit = int(*args.First)
	}

	rs, err := root.db.FindRegions(uint64(limit), uint64(offset))
	if err != nil && err != dbr.ErrNotFound {
		log.Fatal(err)
	}

	connections, err := regions.NewConnectionResolver(root.db, rs, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Types(args arguments.Connection) types.ConnectionResolver {
	offset := 0
	limit := 20
	var err error

	if args.After != nil {
		offset, err = args.After.IntValue()
		if err != nil {
			log.Fatal(err)
		}
	}

	if args.First != nil {
		limit = int(*args.First)
	}

	rs, err := root.db.FindTypes(uint64(limit), uint64(offset))
	if err != nil && err != dbr.ErrNotFound {
		log.Fatal(err)
	}

	connections, err := types.NewConnectionResolver(root.db, rs, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}
