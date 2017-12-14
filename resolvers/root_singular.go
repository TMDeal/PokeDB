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
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&gen, conds); err != nil {
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
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&t, conds); err != nil {
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
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&m, conds); err != nil {
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
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&r, conds); err != nil {
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
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&v, conds); err != nil {
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
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&v, conds); err != nil {
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

	var v models.Ability
	conds := models.NewConditions().Where("id = ?", id).Or("LOWER(name) LIKE LOWER(?)", name)
	if err = root.db.Find(&v, conds); err != nil {
		return nil
	}

	return NewAbilityResolver(root.db, &v)
}
