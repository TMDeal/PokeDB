package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
)

//RootResolver is the root resolver for the graphql schema. All other resolvers
//get returned from this resolver
type RootResolver struct {
	db *models.DB
}

//NewRootResolver returns a new RootResolver
func NewRootResolver(db *models.DB) *RootResolver {
	return &RootResolver{db}
}

func GetSearch(args arguments.Search) (name string, id int, err error) {
	id = -1
	if args.ID != nil {
		_, id, err = FromGlobalID(*args.ID)
		if err != nil {
			return
		}
	}

	if args.Name != nil {
		name = *args.Name + "%"
	}

	return
}

func GetLimitOffset(args arguments.Connection) (int, int, error) {
	offset := 0
	limit := 20

	var err error

	if args.After != nil {
		offset, err = args.After.IntValue()
		if err != nil {
			return 0, 0, err
		}
	}

	if args.First != nil {
		limit = int(*args.First)
	}

	return limit, offset, nil
}
