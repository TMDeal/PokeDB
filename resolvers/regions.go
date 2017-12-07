package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

//RegionResolver resolves the fields of a Region
type RegionResolver struct {
	db     *models.DB
	region *models.Region
}

//NewRegionResolver returns a new RegionResolver
func NewRegionResolver(db *models.DB, r *models.Region) *RegionResolver {
	return &RegionResolver{db, r}
}

//ID resolves the ID field of a Region
func (r *RegionResolver) ID() graphql.ID {
	return GlobalID(models.Region{}, r.region.ID)
}

//Identifier resolves the Identifier field of a Region
func (r *RegionResolver) Identifier() string {
	return r.region.Identifier
}

//Name resolves the Name field of a Region
func (r *RegionResolver) Name() string {
	return r.region.Name
}

//RegionEdgeResolver resolves the fields of an edge
type RegionEdgeResolver struct {
	db     *models.DB
	node   *models.Region
	cursor scalars.Cursor
}

//NewRegionEdgeResolver returns a new RegionEdgeResolver
func NewRegionEdgeResolver(db *models.DB, r *models.Region, c scalars.Cursor) *RegionEdgeResolver {
	return &RegionEdgeResolver{
		db:     db,
		node:   r,
		cursor: c,
	}
}

//Cursor returns the cursor for the edge
func (e *RegionEdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

//Node returns the region for the edge
func (e *RegionEdgeResolver) Node() *RegionResolver {
	return NewRegionResolver(e.db, e.node)
}

//RegionConnectionResolver resolves the fields of a region connection
type RegionConnectionResolver struct {
	db    *models.DB
	items []*models.Region
	start scalars.Cursor
	end   scalars.Cursor
}

//NewRegionConnectionResolver returns a new RegionConnectionResolver
func NewRegionConnectionResolver(db *models.DB, items []*models.Region, args arguments.Connection) (*RegionConnectionResolver, error) {
	start, end, err := MakeCursors(len(items), args)
	if err != nil {
		return nil, err
	}

	return &RegionConnectionResolver{
		db:    db,
		items: items,
		start: *start,
		end:   *end,
	}, nil
}

//TotalCount returns the total number of items in a connection
func (c RegionConnectionResolver) TotalCount() (int32, error) {
	count, err := c.db.Count("regions")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (c RegionConnectionResolver) PageInfo() (*PageResolver, error) {
	count, err := c.TotalCount()
	if err != nil {
		return nil, err
	}
	hasNext, err := HasNextPage(c.end, int(count))
	if err != nil {
		return nil, err
	}

	return NewPageResolver(c.start, c.end, hasNext), nil
}

//Edges returns the edges of a connection
func (c RegionConnectionResolver) Edges() (*[]*RegionEdgeResolver, error) {
	var e []*RegionEdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor(cursorLocation)
		e = append(e, NewRegionEdgeResolver(c.db, item, cursor))
	}

	return &e, nil
}
