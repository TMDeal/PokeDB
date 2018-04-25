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

	var items []models.Move
	query := models.Select("*").From("moves").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
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

	var items []models.Generation
	query := models.Select("*").From("generations").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
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

	var items []models.Region
	query := models.Select("*").From("regions").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
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

	var items []models.Type
	query := models.Select("*").From("types").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
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

	var items []models.Version
	query := models.Select("*").From("versions").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
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

	var items []models.VersionGroup
	query := models.Select("*").From("version_groups").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
		log.Fatal(err)
	}

	connections, err := NewVersionGroupConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Abilities(args arguments.Connection) AbilityConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []models.Ability
	query := models.Select("*").From("abilities").Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
		log.Fatal(err)
	}

	connections, err := NewAbilityConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Items(args arguments.Connection) ItemConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []models.Item
	query := models.Items().Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
		log.Fatal(err)
	}

	connections, err := NewItemConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Machines(args struct {
	arguments.Connection
	VersionGroup int32
}) MachineConnectionResolver {
	limit, offset, err := GetLimitOffset(args.Connection)
	if err != nil {
		log.Fatal(err)
	}

	var items []models.Machine
	query := models.Machines().Limit(limit).Offset(offset).
		Where("version_group_id = ?", args.VersionGroup)

	if err = root.db.FindAll(&items, query); err != nil {
		log.Fatal(err)
	}

	connections, err := NewMachineConnectionResolver(root.db, items, args.Connection)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}

func (root *RootResolver) Berries(args arguments.Connection) BerryConnectionResolver {
	limit, offset, err := GetLimitOffset(args)
	if err != nil {
		log.Fatal(err)
	}

	var items []models.Berry
	query := models.Berries().Limit(limit).Offset(offset)
	if err = root.db.FindAll(&items, query); err != nil {
		log.Fatal(err)
	}

	connections, err := NewBerryConnectionResolver(root.db, items, args)
	if err != nil {
		log.Fatal(err)
	}

	return *connections
}
