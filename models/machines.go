package models

type Machine struct {
	Item
	MachineNumber  int64 `db:"machine_number"`
	VersionGroupID int64 `db:"version_group_id"`
	MoveID         int64 `db:"move_id"`
}

func Machines() *SelectBuilder {
	return Select("*").From("machines AS m").
		Join("items AS i ON i.id = m.item_id")
}

func (m Machine) VersionGroup(f Finder) (*VersionGroup, error) {
	var vg VersionGroup
	query := Select("*").From("version_groups").Where("id = ?", m.VersionGroupID)
	if err := f.Find(&vg, query); err != nil {
		return nil, err
	}

	return &vg, nil
}

func (m Machine) Move(f Finder) (*Move, error) {
	var mv Move
	query := Select("*").From("items").Where("id = ?", m.MoveID)
	if err := f.Find(&mv, query); err != nil {
		return nil, err
	}

	return &mv, nil
}
