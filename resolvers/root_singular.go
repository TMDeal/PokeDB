package resolvers

import (
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
	if err = root.db.Find(&gen, "id = ? OR LOWER(name) LIKE LOWER(?)", id, name); err != nil {
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
	if err = root.db.Find(&t, "id = ? OR LOWER(name) LIKE LOWER(?)", id, name); err != nil {
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
	if err = root.db.Find(&m, "id = ? OR LOWER(name) LIKE LOWER(?)", id, name); err != nil {
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
	if err = root.db.Find(&r, "id = ? OR LOWER(name) LIKE LOWER(?)", id, name); err != nil {
		return nil
	}

	return NewRegionResolver(root.db, &r)
}
