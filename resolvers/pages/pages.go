package pages

import (
	"github.com/TMDeal/PokeDB/scalars"
)

type Resolver struct {
	startCursor *scalars.Cursor
	endCursor   *scalars.Cursor
	hasNextPage bool
}

func NewResolver(start *scalars.Cursor, end *scalars.Cursor, hasNext bool) *Resolver {
	return &Resolver{
		startCursor: start,
		endCursor:   end,
		hasNextPage: hasNext,
	}
}

func (r *Resolver) StartCursor() *scalars.Cursor {
	return r.startCursor
}

func (r *Resolver) EndCursor() *scalars.Cursor {
	return r.endCursor
}

func (r *Resolver) HasNextPage() bool {
	return r.hasNextPage
}
