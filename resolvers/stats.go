package resolvers

//go:generate go run ./connection/main.go -model=Stat -table=stats

import (
	"github.com/TMDeal/PokeDB/models"
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

func (r *StatResolver) GameIndex() int32 {
	return int32(r.stat.GameIndex)
}

func (r *StatResolver) DamageClass() (*DamageClassResolver, error) {
	dc, err := r.stat.DamageClass(r.db)
	if err != nil {
		return nil, err
	}

	return NewDamageClassResolver(r.db, dc), nil
}
