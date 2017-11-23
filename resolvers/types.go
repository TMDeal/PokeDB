package resolvers

import (
	"strconv"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

type TypeEResolver interface {
	EResolver
	Node() *TypeResolver
}

type TypeCResolver interface {
	CResolver
	Edges() (*[]TypeEResolver, error)
}

type TypeResolver struct {
	db *models.DB
	t  *models.Type
}

func NewTypeResolver(db *models.DB, t *models.Type) *TypeResolver {
	return &TypeResolver{db, t}
}

func (r *TypeResolver) ID() graphql.ID {
	id := graphql.ID(strconv.Itoa(int(r.t.ID)))
	return id
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

func (ter *TypeEdgeResolver) Cursor() scalars.Cursor {
	return ter.cursor
}

func (ter *TypeEdgeResolver) Node() *TypeResolver {
	return NewTypeResolver(ter.db, ter.node)
}

type TypeConnectionResolver struct {
	db    *models.DB
	items []*models.Type
	start scalars.Cursor
	end   scalars.Cursor
}

func NewTypeConnectionResolver(db *models.DB, items []*models.Type, args arguments.Connection) (*TypeConnectionResolver, error) {
	start, end, err := MakeCursors(len(items), args)
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
func (tcr TypeConnectionResolver) TotalCount() (int32, error) {
	count, err := tcr.db.Count("types")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (tcr TypeConnectionResolver) PageInfo() (*PageResolver, error) {
	count, err := tcr.TotalCount()
	if err != nil {
		return nil, err
	}
	hasNext, err := HasNextPage(tcr.end, int(count))
	if err != nil {
		return nil, err
	}

	return NewPageResolver(tcr.start, tcr.end, hasNext), nil
}

func (rcr TypeConnectionResolver) Edges() (*[]TypeEResolver, error) {
	var e []TypeEResolver

	for i, item := range rcr.items {
		starti, err := rcr.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor(cursorLocation)
		e = append(e, NewTypeEdgeResolver(rcr.db, item, cursor))
	}

	return &e, nil
}
