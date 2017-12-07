package resolvers

import (
	"fmt"
	"log"

	"github.com/TMDeal/PokeDB/arguments"
)

func (root *RootResolver) Move(args arguments.Search) *MoveResolver {
	if args.ID == nil && args.Name == nil {
		log.Println("NO ARGS")
		return nil
	}

	var err error

	id := -1
	if args.ID != nil {
		_, id, err = FromGlobalID(*args.ID)
		if err != nil {
			log.Println(err)
			return nil
		}
	}

	name := ""
	if args.Name != nil {
		name = *args.Name + "%"
	}
	fmt.Println(name)

	m, err := root.db.FindMove("id = ? OR LOWER(name) LIKE LOWER(?)", int(id), name)
	if err != nil {
		log.Println(err)
		return nil
	}

	return NewMoveResolver(root.db, m)
}

func (root *RootResolver) Region(args arguments.Search) *RegionResolver {
	if args.ID != nil {
		_, id, err := FromGlobalID(*args.ID)
		if err != nil {
			return nil
		}

		r, err := root.db.FindRegion("id = ?", int(id))
		if err != nil {
			return nil
		}

		return NewRegionResolver(root.db, r)
	}

	return nil
}
