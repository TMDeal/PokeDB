package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
)

func (root *RootResolver) Node(args arguments.ID) (*NodeResolver, error) {
	if args.ID != nil {
		typ, id, err := FromGlobalID(*args.ID)
		if err != nil {
			return nil, err
		}

		return ResolveType(root.db, typ, int(id))
	}

	return nil, nil
}

func ResolveType(db *models.DB, t string, id int) (*NodeResolver, error) {
	var node Node

	switch t {
	case "Region":
		var r models.Region
		query := models.Regions().Where("id = ?", id)
		if err := db.Find(&r, query); err != nil {
			return nil, err
		}

		node = NewRegionResolver(db, &r)

	case "Move":
		var m models.Move
		query := models.Moves().Where("id = ?", id)
		if err := db.Find(&m, query); err != nil {
			return nil, err
		}

		node = NewMoveResolver(db, &m)

	case "Generation":
		var gen models.Generation
		query := models.Generations().Where("id = ?", id)
		if err := db.Find(&gen, query); err != nil {
			return nil, err
		}

		node = NewGenerationResolver(db, &gen)

	case "Stat":
		var s models.Stat
		query := models.Stats().Where("id = ?", id)
		if err := db.Find(&s, query); err != nil {
			return nil, err
		}

		node = NewStatResolver(db, &s)

	case "Type":
		var t models.Type
		query := models.Types().Where("id = ?", id)
		if err := db.Find(&t, query); err != nil {
			return nil, err
		}

		node = NewTypeResolver(db, &t)

	case "Version":
		var v models.Version
		query := models.Versions().Where("id = ?", id)
		if err := db.Find(&v, query); err != nil {
			return nil, err
		}

		node = NewVersionResolver(db, &v)

	case "Ability":
		var a models.Ability
		query := models.Abilities().Where("id = ?", id)
		if err := db.Find(&a, query); err != nil {
			return nil, err
		}

		node = NewAbilityResolver(db, &a)

	case "VersionGroup":
		var vg models.VersionGroup
		query := models.VersionGroups().Where("id = ?", id)
		if err := db.Find(&vg, query); err != nil {
			return nil, err
		}

		node = NewVersionGroupResolver(db, &vg)

	case "Item":
		var i models.Item
		query := models.Items().Where("id = ?", id)
		if err := db.Find(&i, query); err != nil {
			return nil, err
		}

		node = NewItemResolver(db, &i)

	case "Machine":
		var m models.Machine
		query := models.Machines().Where("id = ?", id)
		if err := db.Find(&m, query); err != nil {
			return nil, err
		}

		node = NewMachineResolver(db, &m)
	}

	return NewNodeResolver(node), nil
}
