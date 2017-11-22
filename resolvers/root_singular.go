package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/resolvers/generations"
	"github.com/gocraft/dbr"
)

func (root *RootResolver) Generation(args arguments.ID) *generations.Resolver {
	if args.ID != nil {
		gen, err := root.db.FindGeneration("id = ?", int(*args.ID))
		if err == dbr.ErrNotFound {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}

		return generations.NewResolver(root.db, gen)
	}

	return nil
}
