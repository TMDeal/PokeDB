package models

type MoveTarget struct {
	ID          int64  `db:"id"`
	Identifier  string `db:"identifier"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type MoveTargetFinder interface {
	FindMoveTargets(limit uint64, offset uint64) ([]*MoveTarget, error)
	FindMoveTarget(query string, value interface{}) (*MoveTarget, error)
}

func (db DB) FindMoveTargets(limit uint64, offset uint64) ([]*MoveTarget, error) {
	var ms []*MoveTarget
	sess := db.Session()

	_, err := sess.Select("*").
		From("move_targets").
		Limit(limit).
		Offset(offset).
		Load(&ms)

	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (db DB) FindMoveTarget(query string, value interface{}) (*MoveTarget, error) {
	var m MoveTarget
	sess := db.Session()

	_, err := sess.Select("*").
		From("move_targets").
		Where(query, value).
		Load(m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
