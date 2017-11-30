package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/gocraft/dbr"
)

func (root *RootResolver) Node(args arguments.ID) *NodeResolver {
	if args.ID != nil {
		var node Node

		typ, id, err := FromGlobalID(*args.ID)
		if err != nil {
			return nil
		}

		switch typ {
		case "Region":
			r, err := root.db.FindRegion("id = ?", int(id))
			if err == dbr.ErrNotFound {
				return nil
			}
			if err != nil {
				log.Fatal(err)
			}

			node = NewRegionResolver(root.db, r)

		case "Generation":
			gen, err := root.db.FindGeneration("id = ?", int(id))
			if err == dbr.ErrNotFound {
				return nil
			}
			if err != nil {
				log.Fatal(err)
			}

			node = NewGenerationResolver(root.db, gen)

		case "Type":
			t, err := root.db.FindType("id = ?", int(id))
			if err == dbr.ErrNotFound {
				return nil
			}
			if err != nil {
				log.Fatal(err)
			}

			node = NewTypeResolver(root.db, t)
		}

		return NewNodeResolver(node)
	}

	return nil
}
