package resolvers

import (
	"log"

	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/models"
)

func (root *RootResolver) Moves(args arguments.Connection) MoveConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []*models.Move
	if err = root.db.FindAll(&items, models.NewConditions().Limit(limit).Offset(offset)); err != nil {
		log.Fatal(err)
	}

	connections, err := NewMoveConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Generations(args arguments.Connection) GenerationConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []*models.Generation
	if err = root.db.FindAll(&items, models.NewConditions().Limit(limit).Offset(offset)); err != nil {
		log.Fatal(err)
	}

	connections, err := NewGenerationConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Regions(args arguments.Connection) RegionConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []*models.Region
	if err = root.db.FindAll(&items, models.NewConditions().Limit(limit).Offset(offset)); err != nil {
		log.Fatal(err)
	}

	connections, err := NewRegionConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Types(args arguments.Connection) TypeConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []*models.Type
	if err = root.db.FindAll(&items, models.NewConditions().Limit(limit).Offset(offset)); err != nil {
		log.Fatal(err)
	}

	connections, err := NewTypeConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Versions(args arguments.Connection) VersionConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []*models.Version
	if err = root.db.FindAll(&items, models.NewConditions().Limit(limit).Offset(offset)); err != nil {
		log.Fatal(err)
	}

	connections, err := NewVersionConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) VersionGroups(args arguments.Connection) VersionGroupConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []*models.VersionGroup
	if err = root.db.FindAll(&items, models.NewConditions().Limit(limit).Offset(offset)); err != nil {
		log.Fatal(err)
	}

	connections, err := NewVersionGroupConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}
