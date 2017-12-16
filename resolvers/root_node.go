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
			query := models.Select("*").From("regions").Where("id = ?", int(id))
			if err := root.db.Find(&r, query); err != nil {
				return nil
			}

			node = NewRegionResolver(root.db, &r)

		case "Generation":
			var gen models.Generation
			query := models.Select("*").From("generations").Where("id = ?", int(id))
			if err := root.db.Find(&gen, query); err != nil {
				return nil
			}

			node = NewGenerationResolver(root.db, &gen)

		case "Type":
			var t models.Type
			query := models.Select("*").From("types").Where("id = ?", int(id))
			if err := root.db.Find(&t, query); err != nil {
				return nil
			}

			node = NewTypeResolver(root.db, &t)

		case "Move":
			var m models.Move
			query := models.Select("*").From("moves").Where("id = ?", int(id))
			if err := root.db.Find(&m, query); err != nil {
				return nil
			}

			node = NewMoveResolver(root.db, &m)
		}

		return NewNodeResolver(node)
	}

	return nil
}
