package resolvers

//go:generate ../connection -model=Move -table=moves

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
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

func (m *MoveResolver) Flags() *[]*MoveFlagResolver {
	var mfs []*MoveFlagResolver

	mf, err := m.move.Flags(m.db)
	if err != nil {
		return nil
	}

	for _, v := range mf {
		mfs = append(mfs, NewMoveFlagResolver(m.db, &v))
	}

	return &mfs
}

func (m *MoveResolver) FlavorText(args arguments.FlavorText) (*FlavorTextResolver, error) {
	flav, err := m.move.FlavorText(m.db, int(args.VersionGroup))
	if err != nil {
		return nil, err
	}

	return NewFlavorTextResolver(m.db, flav), nil
}

func (m *MoveResolver) Effect() (*MoveEffectResolver, error) {
	me, err := m.move.Effect(m.db)
	if err != nil {
		return nil, err
	}

	return NewMoveEffectResolver(m.db, me), nil
}

func (m *MoveResolver) Target() (*MoveTargetResolver, error) {
	mt, err := m.move.Target(m.db)
	if err != nil {
		return nil, err
	}

	return NewMoveTargetResolver(m.db, mt), nil
}

func (m *MoveResolver) Meta() (*MoveMetaResolver, error) {
	meta, err := m.move.Meta(m.db)
	if err != nil {
		return nil, err
	}

	return NewMoveMetaResolver(m.db, meta), nil
}

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
