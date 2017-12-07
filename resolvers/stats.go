package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

type StatResolver struct {
	db   *models.DB
	stat *models.Stat
}

func NewStatResolver(db *models.DB, s *models.Stat) *StatResolver {
	return &StatResolver{db, s}
}

func (r *StatResolver) ID() graphql.ID {
	return GlobalID(models.Stat{}, r.stat.ID)
}

func (r *StatResolver) Identifier() string {
	return r.stat.Identifier
}

func (r *StatResolver) Name() string {
	return r.stat.Name
}

func (r *StatResolver) BattleOnly() bool {
	return r.stat.BattleOnly
}

func (r *StatResolver) GameIndex() int {
	return r.stat.GameIndex
}

func (r *StatResolver) DamageClass() (*DamageClassResolver, error) {
	dc, err := r.stat.DamageClass(r.db)
	if err != nil {
		return nil, err
	}

	return NewDamageClassResolver(r.db, dc), nil
}

type StatEdgeResolver struct {
	db     *models.DB
	node   *models.Stat
	cursor scalars.Cursor
}

func NewStatEdgeResolver(db *models.DB, gen *models.Stat, c scalars.Cursor) *StatEdgeResolver {
	return &StatEdgeResolver{
		db:     db,
		node:   gen,
		cursor: c,
	}
}

func (e *StatEdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

func (e *StatEdgeResolver) Node() *StatResolver {
	return NewStatResolver(e.db, e.node)
}

type StatConnectionResolver struct {
	db    *models.DB
	items []*models.Stat
	start scalars.Cursor
	end   scalars.Cursor
}

func NewStatConnectionResolver(db *models.DB, items []*models.Stat, args arguments.Connection) (*StatConnectionResolver, error) {
	start, end, err := MakeCursors(len(items), args)
	if err != nil {
		return nil, err
	}

	return &StatConnectionResolver{
		db:    db,
		items: items,
		start: *start,
		end:   *end,
	}, nil
}

//TotalCount returns the total number of items in a connection
func (c StatConnectionResolver) TotalCount() (int32, error) {
	count, err := c.db.Count("stats")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (c StatConnectionResolver) PageInfo() (*PageResolver, error) {
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

func (c StatConnectionResolver) Edges() (*[]*StatEdgeResolver, error) {
	var e []*StatEdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor(cursorLocation)
		e = append(e, NewStatEdgeResolver(c.db, item, cursor))
	}

	return &e, nil
}
