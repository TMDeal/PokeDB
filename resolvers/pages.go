package resolvers

import (
	"github.com/TMDeal/PokeDB/scalars"
)

//PageResolver resolves the fields for pageInfo
type PageResolver struct {
	startCursor scalars.Cursor
	endCursor   scalars.Cursor
	hasNextPage bool
}

//NewPageResolver returns a new PageResolver
func NewPageResolver(start scalars.Cursor, end scalars.Cursor, hasNext bool) *PageResolver {
	return &PageResolver{
		startCursor: start,
		endCursor:   end,
		hasNextPage: hasNext,
	}
}

//StartCursor returns the starting cursor
func (r *PageResolver) StartCursor() scalars.Cursor {
	return r.startCursor
}

//EndCursor returns the ending cursor
func (r *PageResolver) EndCursor() scalars.Cursor {
	return r.endCursor
}

//hasNextPage returns true or false depending on if there is a next page
func (r *PageResolver) HasNextPage() bool {
	return r.hasNextPage
}
