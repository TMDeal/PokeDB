package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/gocraft/dbr"
)

func (root *RootResolver) Region(args arguments.ID) *RegionResolver {
	if args.ID != nil {
		r, err := root.db.FindRegion("id = ?", int(*args.ID))
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
