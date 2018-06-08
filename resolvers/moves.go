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

func (r MoveResolver) ID() graphql.ID {
	return GlobalID(models.Move{}, r.move.ID)
}

func (r MoveResolver) Identifier() string {
	return r.move.Identifier
}

func (r MoveResolver) Name() string {
	return r.move.Name
}

func (r MoveResolver) ContestEffect() *ContestEffectResolver {
	c, err := r.move.ContestEffect(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}
	return NewContestEffectResolver(r.db, c)
}

func (r MoveResolver) ContestType() *ContestTypeResolver {
	c, err := r.move.ContestType(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}
	return NewContestTypeResolver(r.db, c)
}

func (r MoveResolver) SuperContestEffect() *SuperContestEffectResolver {
	c, err := r.move.SuperContestEffect(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}
	return NewSuperContestEffectResolver(r.db, c)
}

func (r MoveResolver) Generation() *GenerationResolver {
	gen, err := r.move.Generation(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}
	return NewGenerationResolver(r.db, gen)
}

func (r MoveResolver) DamageClass() *DamageClassResolver {
	dc, err := r.move.DamageClass(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}
	return NewDamageClassResolver(r.db, dc)
}

func (r MoveResolver) Type() *TypeResolver {
	t, err := r.move.Type(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}
	return NewTypeResolver(r.db, t)
}

func (r MoveResolver) Flags() *[]*MoveFlagResolver {
	mf, err := r.move.Flags(r.db)
	if err != nil {
		r.db.Log(err)
		return nil
	}

	var mfs []*MoveFlagResolver
	for i, _ := range mf {
		mfs = append(mfs, NewMoveFlagResolver(r.db, &mf[i]))
	}

	return &mfs
}

func (r MoveResolver) FlavorText(args arguments.FlavorText) (*FlavorTextResolver, error) {
	flav, err := r.move.FlavorText(r.db, int(args.VersionGroup))
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewFlavorTextResolver(r.db, flav), nil
}

func (r MoveResolver) Effect() (*MoveEffectResolver, error) {
	me, err := r.move.Effect(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewMoveEffectResolver(r.db, me), nil
}

func (r MoveResolver) Target() (*MoveTargetResolver, error) {
	mt, err := r.move.Target(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewMoveTargetResolver(r.db, mt), nil
}

func (r MoveResolver) Meta() (*MoveMetaResolver, error) {
	meta, err := r.move.Meta(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewMoveMetaResolver(r.db, meta), nil
}

func (r MoveResolver) Power() *int32 {
	if !r.move.Power.Valid {
		return nil
	}
	p := int32(r.move.Power.Int64)
	return &p
}

func (r MoveResolver) PP() *int32 {
	if !r.move.PP.Valid {
		return nil
	}
	pp := int32(r.move.PP.Int64)
	return &pp
}

func (r MoveResolver) Accuracy() *int32 {
	if !r.move.Accuracy.Valid {
		return nil
	}
	acc := int32(r.move.Accuracy.Int64)
	return &acc
}

func (r MoveResolver) Priority() int32 {
	return int32(r.move.PP.Int64)
}
