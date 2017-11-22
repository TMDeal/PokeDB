package regions

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/resolvers/pages"
	"github.com/TMDeal/PokeDB/scalars"
)

type ConnectionResolver struct {
	db    *models.DB
	items []*models.Region
	start scalars.Cursor
	end   scalars.Cursor
}

func NewConnectionResolver(db *models.DB, items []*models.Region, args arguments.Connection) (*ConnectionResolver, error) {
	start := scalars.NewCursor(0)
	if args.After != nil {
		start = *args.After
	}

	starti, err := start.IntValue()
	if err != nil {
		return nil, err
	}

	end := scalars.NewCursor(len(items) + starti)
	if args.First != nil {
		actualEnd := starti + int(*args.First)
		if actualEnd > len(items) {
			end = scalars.NewCursor(actualEnd)
		}
	}

	return &ConnectionResolver{
		db:    db,
		items: items,
		start: start,
		end:   end,
	}, nil
}

func (c ConnectionResolver) TotalCount() int32 {
	return int32(len(c.items))
}

func (c ConnectionResolver) PageInfo() (*pages.Resolver, error) {
	starti, err := c.start.IntValue()
	if err != nil {
		return nil, err
	}

	hasNext := starti < len(c.items)

	return pages.NewResolver(&c.start, &c.end, hasNext), nil
}

func (c ConnectionResolver) Edges() (*[]*EdgeResolver, error) {
	var e []*EdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		e = append(e, NewEdgeResolver(NewResolver(c.db, item), scalars.NewCursor(cursorLocation)))
	}

	return &e, nil
}
