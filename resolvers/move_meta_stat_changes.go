package resolvers

import "github.com/TMDeal/PokeDB/models"

type MoveMetaStatChangeResolver struct {
	db  *models.DB
	msc *models.MoveMetaStatChange
}

func NewMoveMetaStatChangeResolver(db *models.DB, msc *models.MoveMetaStatChange) *MoveMetaStatChangeResolver {
	return &MoveMetaStatChangeResolver{db, msc}
}

func (r MoveMetaStatChangeResolver) Change() int32 {
	return int32(r.msc.Change)
}

func (r MoveMetaStatChangeResolver) Stat() (*StatResolver, error) {
	s, err := r.msc.Stat(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewStatResolver(r.db, s), nil
}
