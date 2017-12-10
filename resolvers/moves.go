package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

type MoveResolver struct {
	db   *models.DB
	move *models.Move
}

func NewMoveResolver(db *models.DB, m *models.Move) *MoveResolver {
	return &MoveResolver{db, m}
}

func (m *MoveResolver) ID() graphql.ID {
	return GlobalID(models.Move{}, m.move.ID)
}

func (m *MoveResolver) Identifier() string {
	return m.move.Identifier
}

func (m *MoveResolver) Name() string {
	return m.move.Name
}

func (m *MoveResolver) ContestEffect() *ContestEffectResolver {
	c, err := m.move.ContestEffect(m.db)
	if err != nil {
		return nil
	}
	return NewContestEffectResolver(m.db, c)
}

func (m *MoveResolver) ContestType() *ContestTypeResolver {
	c, err := m.move.ContestType(m.db)
	if err != nil {
		return nil
	}
	return NewContestTypeResolver(m.db, c)
}

func (m *MoveResolver) SuperContestEffect() *SuperContestEffectResolver {
	c, err := m.move.SuperContestEffect(m.db)
	if err != nil {
		return nil
	}
	return NewSuperContestEffectResolver(m.db, c)
}

func (m *MoveResolver) Generation() *GenerationResolver {
	gen, err := m.move.Generation(m.db)
	if err != nil {
		return nil
	}
	return NewGenerationResolver(m.db, gen)
}

func (m *MoveResolver) DamageClass() *DamageClassResolver {
	dc, err := m.move.DamageClass(m.db)
	if err != nil {
		return nil
	}
	return NewDamageClassResolver(m.db, dc)
}

func (m *MoveResolver) Type() *TypeResolver {
	t, err := m.move.Type(m.db)
	if err != nil {
		return nil
	}
	return NewTypeResolver(m.db, t)
}

// func (m *MoveResolver) Target() *MoveTargetResolver {}

func (m *MoveResolver) Power() *int32 {
	if !m.move.Power.Valid {
		return nil
	}
	p := int32(m.move.Power.Int64)
	return &p
}

func (m *MoveResolver) PP() *int32 {
	if !m.move.PP.Valid {
		return nil
	}
	pp := int32(m.move.PP.Int64)
	return &pp
}

func (m *MoveResolver) Accuracy() *int32 {
	if !m.move.Accuracy.Valid {
		return nil
	}
	acc := int32(m.move.Accuracy.Int64)
	return &acc
}

func (m *MoveResolver) Priority() int32 {
	return int32(m.move.PP.Int64)
}

type MoveEdgeResolver struct {
	db     *models.DB
	node   *models.Move
	cursor scalars.Cursor
}

func NewMoveEdgeResolver(db *models.DB, gen *models.Move, c scalars.Cursor) *MoveEdgeResolver {
	return &MoveEdgeResolver{
		db:     db,
		node:   gen,
		cursor: c,
	}
}

func (e *MoveEdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

func (e *MoveEdgeResolver) Node() *MoveResolver {
	return NewMoveResolver(e.db, e.node)
}

type MoveConnectionResolver struct {
	db    *models.DB
	items []*models.Move
	start scalars.Cursor
	end   scalars.Cursor
}

func NewMoveConnectionResolver(db *models.DB, items []*models.Move, args arguments.Connection) (*MoveConnectionResolver, error) {
	start, end, err := MakeCursors("regions", len(items), args)
	if err != nil {
		return nil, err
	}

	return &MoveConnectionResolver{
		db:    db,
		items: items,
		start: *start,
		end:   *end,
	}, nil
}

//TotalCount returns the total number of items in a connection
func (c MoveConnectionResolver) TotalCount() (int32, error) {
	count, err := c.db.Count("moves")
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

//PageInfo returns the information about the current page
func (c MoveConnectionResolver) PageInfo() (*PageResolver, error) {
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

func (c MoveConnectionResolver) Edges() (*[]*MoveEdgeResolver, error) {
	var e []*MoveEdgeResolver

	for i, item := range c.items {
		starti, err := c.start.IntValue()
		if err != nil {
			return nil, err
		}
		cursorLocation := starti + i + 1
		cursor := scalars.NewCursor("regions", cursorLocation)
		e = append(e, NewMoveEdgeResolver(c.db, item, cursor))
	}

	return &e, nil
}
