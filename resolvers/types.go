package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

type TypeResolver struct {
	db *models.DB
	t  *models.Type
}

func NewTypeResolver(db *models.DB, t *models.Type) *TypeResolver {
	return &TypeResolver{db, t}
}

func (r *TypeResolver) ID() graphql.ID {
	return GlobalID(models.Type{}, r.t.ID)
}

func (r *TypeResolver) Identifier() string {
	return r.t.Identifier
}

func (r *TypeResolver) Name() string {
	return r.t.Name
}

func (r *TypeResolver) Generation() *GenerationResolver {
	gen, err := r.t.Generation(r.db)
	if err != nil {
		return nil
	}

	return NewGenerationResolver(r.db, gen)
}

func (r *TypeResolver) DamageClass() *DamageClassResolver {
	dc, err := r.t.DamageClass(r.db)
	if err != nil {
		return nil
	}

	return NewDamageClassResolver(r.db, dc)
}

type TypeEdgeResolver struct {
	db     *models.DB
	node   *models.Type
	cursor scalars.Cursor
}

func NewTypeEdgeResolver(db *models.DB, gen *models.Type, c scalars.Cursor) *TypeEdgeResolver {
	return &TypeEdgeResolver{
		db:     db,
		node:   gen,
		cursor: c,
	}
}

func (e *TypeEdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

func (e *TypeEdgeResolver) Node() *TypeResolver {
	return NewTypeResolver(e.db, e.node)
}

type TypeConnectionResolver struct {
	db    *models.DB
	items []*models.Type
	start scalars.Cursor
	end   scalars.Cursor
}

func NewTypeConnectionResolver(db *models.DB, items []*models.Type, args arguments.Connection) (*TypeConnectionResolver, error) {
	start, end, err := MakeCursors("types", len(items), args)
	if err != nil {
		return nil, err
	}

	return &TypeConnectionResolver{
		db:    db,
		items: items,
		start: *start,
		end:   *end,
	}, nil
}

//TotalCount returns the total number of items in a connection
func (c TypeConnectionResolver) TotalCount() (int32, error) {
	count, err := c.db.Count("types")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (c TypeConnectionResolver) PageInfo() (*PageResolver, error) {
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

func (c TypeConnectionResolver) Edges() (*[]*TypeEdgeResolver, error) {
	var e []*TypeEdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor("types", cursorLocation)
		e = append(e, NewTypeEdgeResolver(c.db, item, cursor))
	}

	return &e, nil
}
