package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
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
			if err != nil {
				return nil
			}

			node = NewRegionResolver(root.db, r)

		case "Generation":
			gen, err := root.db.FindGeneration("id = ?", int(id))
			if err != nil {
				return nil
			}

			node = NewGenerationResolver(root.db, gen)

		case "Type":
			t, err := root.db.FindType("id = ?", int(id))
			if err != nil {
				return nil
			}

			node = NewTypeResolver(root.db, t)

		case "Move":
			m, err := root.db.FindMove("id = ?", int(id))
			if err != nil {
				return nil
			}

			node = NewMoveResolver(root.db, m)
		}

		return NewNodeResolver(node)
	}

	return nil
}
