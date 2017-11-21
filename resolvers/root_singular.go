package resolvers

import (
	"database/sql"
	"log"

	"github.com/TMDeal/PokeDB/resolvers/generations"
)

func (root *RootResolver) Generation(args struct{ ID int32 }) *generations.Resolver {
	gen, err := root.db.FindGeneration("id = ?", int(args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	return generations.NewResolver(root.db, gen)
}
