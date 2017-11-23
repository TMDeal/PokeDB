package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

type GenerationEResolver interface {
	EResolver
	Node() *GenerationResolver
}

type GenerationCResolver interface {
	CResolver
	Edges() (*[]GenerationEResolver, error)
}

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
	id := graphql.ID(strconv.Itoa(int(r.generation.ID)))
	return id
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

func (ger *GenerationEdgeResolver) Cursor() scalars.Cursor {
	return ger.cursor
}

func (ger *GenerationEdgeResolver) Node() *GenerationResolver {
	return NewGenerationResolver(ger.db, ger.node)
}

type GenerationConnectionResolver struct {
	db    *models.DB
	items []*models.Generation
	start scalars.Cursor
	end   scalars.Cursor
}

func NewGenerationConnectionResolver(db *models.DB, items []*models.Generation, args arguments.Connection) (*GenerationConnectionResolver, error) {
	start, end, err := MakeCursors(len(items), args)
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
func (gcr GenerationConnectionResolver) TotalCount() (int32, error) {
	count, err := gcr.db.Count("generations")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (gcr GenerationConnectionResolver) PageInfo() (*PageResolver, error) {
	count, err := gcr.TotalCount()
	if err != nil {
		return nil, err
	}
	hasNext, err := HasNextPage(gcr.end, int(count))
	if err != nil {
		return nil, err
	}

	return NewPageResolver(gcr.start, gcr.end, hasNext), nil
}

func (rcr GenerationConnectionResolver) Edges() (*[]GenerationEResolver, error) {
	var e []GenerationEResolver

	for i, item := range rcr.items {
		starti, err := rcr.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor(cursorLocation)
		e = append(e, NewGenerationEdgeResolver(rcr.db, item, cursor))
	}

	return &e, nil
}
