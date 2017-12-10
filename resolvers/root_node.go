package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
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
			var r models.Region
			if err := root.db.Find(&r, models.NewConditions().Where("id = ?", int(id))); err != nil {
				return nil
			}

			node = NewRegionResolver(root.db, &r)

		case "Generation":
			var gen models.Generation
			if err := root.db.Find(&gen, models.NewConditions().Where("id = ?", int(id))); err != nil {
				return nil
			}

			node = NewGenerationResolver(root.db, &gen)

		case "Type":
			var t models.Type
			if err := root.db.Find(&t, models.NewConditions().Where("id = ?", int(id))); err != nil {
				return nil
			}

			node = NewTypeResolver(root.db, &t)

		case "Move":
			var m models.Move
			if err := root.db.Find(&m, models.NewConditions().Where("id = ?", int(id))); err != nil {
				return nil
			}

			node = NewMoveResolver(root.db, &m)
		}

		return NewNodeResolver(node)
	}

	return nil
}
