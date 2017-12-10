package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

//GenerationResolver resolves the fields of a Generation
type GenerationResolver struct {
	db         *models.DB
	generation *models.Generation
}

//NewGenerationResolver returns a new Resolver
func NewGenerationResolver(db *models.DB, g *models.Generation) *GenerationResolver {
	return &GenerationResolver{db, g}
}

//ID resolves the ID field of a Generation
func (r *GenerationResolver) ID() graphql.ID {
	return GlobalID(models.Generation{}, r.generation.ID)
}

//Identifier resolves the Identifier field of a Generation
func (r *GenerationResolver) Identifier() string {
	return r.generation.Identifier
}

//Name resolves the Name field of a Generation
func (r *GenerationResolver) Name() string {
	return r.generation.Name
}

//Region resolves the Region of a generation, by finding the region based on a
//Generations RegionID, and returning a RegionResolver for that Region
func (r *GenerationResolver) Region() *RegionResolver {
	rn, err := r.generation.Region(r.db)
	if err != nil {
		return nil
	}
	return NewRegionResolver(r.db, rn)
}

type GenerationEdgeResolver struct {
	db     *models.DB
	node   *models.Generation
	cursor scalars.Cursor
}

func NewGenerationEdgeResolver(db *models.DB, gen *models.Generation, c scalars.Cursor) *GenerationEdgeResolver {
	return &GenerationEdgeResolver{
		db:     db,
		node:   gen,
		cursor: c,
	}
}

func (e *GenerationEdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

func (e *GenerationEdgeResolver) Node() *GenerationResolver {
	return NewGenerationResolver(e.db, e.node)
}

type GenerationConnectionResolver struct {
	db    *models.DB
	items []*models.Generation
	start scalars.Cursor
	end   scalars.Cursor
}

func NewGenerationConnectionResolver(db *models.DB, items []*models.Generation, args arguments.Connection) (*GenerationConnectionResolver, error) {
	start, end, err := MakeCursors("generations", len(items), args)
	if err != nil {
		return nil, err
	}

	return &GenerationConnectionResolver{
		db:    db,
		items: items,
		start: *start,
		end:   *end,
	}, nil
}

//TotalCount returns the total number of items in a connection
func (c GenerationConnectionResolver) TotalCount() (int32, error) {
	count, err := c.db.Count("generations")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (c GenerationConnectionResolver) PageInfo() (*PageResolver, error) {
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

func (c GenerationConnectionResolver) Edges() (*[]*GenerationEdgeResolver, error) {
	var e []*GenerationEdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor("generations", cursorLocation)
		e = append(e, NewGenerationEdgeResolver(c.db, item, cursor))
	}

	return &e, nil
}
