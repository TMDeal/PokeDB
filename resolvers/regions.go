package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

//RegionEResolver is an interface for an edge in a region connection
type RegionEResolver interface {
	EResolver
	Node() *RegionResolver
}

//RegionCResolver is an interface for a connection that returns info for
//regions in a pagination friendly manner
type RegionCResolver interface {
	CResolver
	Edges() (*[]RegionEResolver, error)
}

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
	id := graphql.ID(strconv.Itoa(int(r.region.ID)))
	return id
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
func (rer *RegionEdgeResolver) Cursor() scalars.Cursor {
	return rer.cursor
}

//Node returns the region for the edge
func (rer *RegionEdgeResolver) Node() *RegionResolver {
	return NewRegionResolver(rer.db, rer.node)
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
func (rcr RegionConnectionResolver) TotalCount() (int32, error) {
	count, err := rcr.db.Count("regions")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (rcr RegionConnectionResolver) PageInfo() (*PageResolver, error) {
	count, err := rcr.TotalCount()
	if err != nil {
		return nil, err
	}
	hasNext, err := HasNextPage(rcr.end, int(count))
	if err != nil {
		return nil, err
	}

	return NewPageResolver(rcr.start, rcr.end, hasNext), nil
}

//Edges returns the edges of a connection
func (rcr RegionConnectionResolver) Edges() (*[]RegionEResolver, error) {
	var e []RegionEResolver

	for i, item := range rcr.items {
		starti, err := rcr.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor(cursorLocation)
		e = append(e, NewRegionEdgeResolver(rcr.db, item, cursor))
	}

	return &e, nil
}
