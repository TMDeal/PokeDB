package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/gocraft/dbr"
)

func (root *RootResolver) Region(args arguments.Search) *RegionResolver {
	if args.ID != nil {
		_, id, err := FromGlobalID(*args.ID)
		if err != nil {
			return nil
		}

		r, err := root.db.FindRegion("id = ?", int(id))
		if err == dbr.ErrNotFound {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}

		return NewRegionResolver(root.db, r)
	}

	return nil
}
