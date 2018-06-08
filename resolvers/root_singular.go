package resolvers

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
)

func (root *RootResolver) Generation(args arguments.Search) *GenerationResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var gen models.Generation
	query := sq.Select("*").From("generations").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&gen, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewGenerationResolver(root.db, &gen)
}

func (root *RootResolver) Type(args arguments.Search) *TypeResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var t models.Type
	query := sq.Select("*").From("types").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&t, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewTypeResolver(root.db, &t)
}

func (root *RootResolver) Move(args arguments.Search) *MoveResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var m models.Move
	query := sq.Select("*").From("moves").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&m, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewMoveResolver(root.db, &m)
}

func (root *RootResolver) Region(args arguments.Search) *RegionResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var r models.Region
	query := sq.Select("*").From("regions").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&r, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewRegionResolver(root.db, &r)
}

func (root *RootResolver) Version(args arguments.Search) *VersionResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var v models.Version
	query := sq.Select("*").From("versions").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&v, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewVersionResolver(root.db, &v)
}

func (root *RootResolver) VersionGroup(args arguments.Search) *VersionGroupResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var v models.VersionGroup
	query := sq.Select("*").From("version_groups").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&v, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewVersionGroupResolver(root.db, &v)
}

func (root *RootResolver) Ability(args arguments.Search) *AbilityResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var a models.Ability
	query := sq.Select("*").From("abilities").Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&a, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewAbilityResolver(root.db, &a)
}

func (root *RootResolver) Item(args arguments.Search) *ItemResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var a models.Item
	query := models.Items().Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&a, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewItemResolver(root.db, &a)
}

func (root *RootResolver) Machine(args struct {
	arguments.Search
	VersionGroup int32
}) *MachineResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args.Search)
	if err != nil {
		root.db.Log(err)
		return nil
	}

	var a models.Machine
	query := models.Machines().Where("id = ? OR LOWER(name) LIKE LOWER(?) AND version_group_id = ?", id, name, args.VersionGroup)

	if err = root.db.Find(&a, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewMachineResolver(root.db, &a)
}

func (root *RootResolver) Berry(args arguments.Search) *BerryResolver {
	if args.ID == nil && args.Name == nil {
		return nil
	}

	name, id, err := GetSearch(args)
	if err != nil {
		return nil
	}

	var a models.Berry
	query := models.Berries().Where("id = ? OR LOWER(name) LIKE LOWER(?)", id, name)
	if err = root.db.Find(&a, query); err != nil {
		root.db.Log(err)
		return nil
	}

	return NewBerryResolver(root.db, &a)
}
