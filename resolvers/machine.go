package resolvers

//go:generate ../connection -model=Machine -table=machines

import (
	"github.com/TMDeal/PokeDB/models"
)

type MachineResolver struct {
	*ItemResolver
	db *models.DB
	m  *models.Machine
}

func NewMachineResolver(db *models.DB, m *models.Machine) *MachineResolver {
	return &MachineResolver{NewItemResolver(db, &m.Item), db, m}
}

func (r MachineResolver) Number() int32 {
	return int32(r.m.MachineNumber)
}

func (r MachineResolver) VersionGroup() (*VersionGroupResolver, error) {
	vg, err := r.m.VersionGroup(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewVersionGroupResolver(r.db, vg), nil
}

func (r MachineResolver) Move() (*MoveResolver, error) {
	mv, err := r.m.Move(r.db)
	if err != nil {
		r.db.Log(err)
		return nil, err
	}

	return NewMoveResolver(r.db, mv), nil
}
