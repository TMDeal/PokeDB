package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveMetaResolver struct {
	db   *models.DB
	meta *models.MoveMeta
}

func NewMoveMetaResolver(db *models.DB, meta *models.MoveMeta) *MoveMetaResolver {
	return &MoveMetaResolver{db, meta}
}

func (r MoveMetaResolver) MinHits() *int32 {
	if !r.meta.MinHits.Valid {
		return nil
	}

	mh := int32(r.meta.MinHits.Int64)
	return &mh
}

func (r MoveMetaResolver) MaxHits() *int32 {
	if !r.meta.MaxHits.Valid {
		return nil
	}

	mh := int32(r.meta.MaxHits.Int64)
	return &mh
}

func (r MoveMetaResolver) MinTurns() *int32 {
	if !r.meta.MaxHits.Valid {
		return nil
	}

	mt := int32(r.meta.MinTurns.Int64)
	return &mt
}

func (r MoveMetaResolver) MaxTurns() *int32 {
	if !r.meta.MaxHits.Valid {
		return nil
	}

	mt := int32(r.meta.MaxTurns.Int64)
	return &mt
}

func (r MoveMetaResolver) Drain() int32 {
	return int32(r.meta.Drain)
}

func (r MoveMetaResolver) Healing() int32 {
	return int32(r.meta.Healing)
}

func (r MoveMetaResolver) CritRate() int32 {
	return int32(r.meta.CritRate)
}

func (r MoveMetaResolver) AilmentChance() int32 {
	return int32(r.meta.AilmentChance)
}

func (r MoveMetaResolver) FlinchChance() int32 {
	return int32(r.meta.FlinchChance)
}

func (r MoveMetaResolver) StatChance() int32 {
	return int32(r.meta.StatChance)
}

func (r MoveMetaResolver) Ailment() (*MoveMetaAilmentResolver, error) {
	ail, err := r.meta.Ailment(r.db)
	if err != nil {
		return nil, err
	}

	return NewMoveMetaAilmentResolver(r.db, ail), nil
}

func (r MoveMetaResolver) Category() (*MoveMetaCategoryResolver, error) {
	cat, err := r.meta.Category(r.db)
	if err != nil {
		return nil, err
	}

	return NewMoveMetaCategoryResolver(r.db, cat), nil
}

func (r MoveMetaResolver) StatChanges() ([]*MoveMetaStatChangeResolver, error) {
	msc, err := r.meta.StatChanges(r.db)
	if err != nil {
		return nil, err
	}

	var mscr []*MoveMetaStatChangeResolver

	for _, v := range msc {
		mscr = append(mscr, NewMoveMetaStatChangeResolver(r.db, &v))
	}

	return mscr, nil
}
