package resolvers

import (
	"log"

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
	query := models.Select("*").From("generations").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&gen, query); err != nil {
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
	query := models.Select("*").From("types").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&t, query); err != nil {
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
	query := models.Select("*").From("moves").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&m, query); err != nil {
		log.Println(err)
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
	query := models.Select("*").From("regions").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&r, query); err != nil {
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
	query := models.Select("*").From("versions").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&v, query); err != nil {
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
	query := models.Select("*").From("version_groups").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&v, query); err != nil {
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
	query := models.Select("*").From("abilities").Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&a, query); err != nil {
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
	query := models.Items().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&a, query); err != nil {
		return nil
	}

	return NewItemResolver(root.db, &a)
}
