package resolvers

import "github.com/TMDeal/PokeDB/models"

type TypeEfficacyResolver struct {
	db *models.DB
	t  *models.Type
}

type TypeEfficacyDirectionResolver struct {
	db   *models.DB
	t    *models.Type
	mult models.DamageMult
}

func NewTypeEfficacyResolver(db *models.DB, t *models.Type) *TypeEfficacyResolver {
	return &TypeEfficacyResolver{db, t}
}

func NewTypeEfficacyDirectionResolver(db *models.DB, t *models.Type, mult models.DamageMult) *TypeEfficacyDirectionResolver {
	return &TypeEfficacyDirectionResolver{db, t, mult}
}

func (r TypeEfficacyResolver) Double() *TypeEfficacyDirectionResolver {
	return NewTypeEfficacyDirectionResolver(r.db, r.t, models.DoubleDamage)
}

func (r TypeEfficacyResolver) Normal() *TypeEfficacyDirectionResolver {
	return NewTypeEfficacyDirectionResolver(r.db, r.t, models.NormalDamage)
}

func (r TypeEfficacyResolver) Half() *TypeEfficacyDirectionResolver {
	return NewTypeEfficacyDirectionResolver(r.db, r.t, models.HalfDamage)
}

func (r TypeEfficacyResolver) None() *TypeEfficacyDirectionResolver {
	return NewTypeEfficacyDirectionResolver(r.db, r.t, models.NoDamage)
}

func (r TypeEfficacyDirectionResolver) From() ([]*TypeResolver, error) {
	ts, err := r.t.DamageFrom(r.db, r.mult)
	if err != nil {
		return nil, err
	}

	var trs []*TypeResolver
	for _, t := range ts {
		trs = append(trs, NewTypeResolver(r.db, t))
	}

	return trs, nil
}

func (r TypeEfficacyDirectionResolver) To() ([]*TypeResolver, error) {
	ts, err := r.t.DamageTo(r.db, r.mult)
	if err != nil {
		return nil, err
	}

	var trs []*TypeResolver
	for _, t := range ts {
		trs = append(trs, NewTypeResolver(r.db, t))
	}

	return trs, nil
}
