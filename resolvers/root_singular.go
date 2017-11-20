package resolvers

import (
	"database/sql"
	"log"
)

func (root *RootResolver) Generation(args struct{ ID int32 }) *GenerationResolver {
	gen, err := root.db.FindGeneration(int(args.ID))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	return NewGenerationResolver(root.db, gen)
}
