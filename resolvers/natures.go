package resolvers

//go:generate ../connection -model=Nature -table=natures

import (
	"github.com/TMDeal/PokeDB/models"
	graphql "github.com/neelance/graphql-go"
)

type NatureResolver struct {
	db *models.DB
	n  *models.Nature
}

func NewNatureResolver(db *models.DB, n *models.Nature) *NatureResolver {
	return &NatureResolver{db, n}
}

func (r NatureResolver) ID() graphql.ID {
	return GlobalID(models.Nature{}, r.n.ID)
}

func (r NatureResolver) Identifier() string {
	return r.n.Identifier
}

func (r NatureResolver) Name() string {
	return r.n.Name
}

func (r NatureResolver) DecreasedStat() (*StatResolver, error) {
	stat, err := r.n.Decreased(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewStatResolver(r.db, stat), nil
}

func (r NatureResolver) IncreasedStat() (*StatResolver, error) {
	stat, err := r.n.Increased(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewStatResolver(r.db, stat), nil
}

func (r NatureResolver) LikesFlavor() (string, error) {
	flav, err := r.n.Likes(r.db)
	if err != nil {
		r.db.Log(err)
		return "", err
	}

	return flav.Flavor, nil
	// return NewContestTypeResolver(r.db, flav), nil
}

func (r NatureResolver) HatesFlavor() (string, error) {
	flav, err := r.n.Hates(r.db)
	if err != nil {
		r.db.Log(err)
		return "", err
	}

	return flav.Flavor, nil
	// return NewContestTypeResolver(r.db, flav), nil
}
