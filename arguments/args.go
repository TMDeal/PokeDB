package arguments

import (
	"github.com/TMDeal/PokeDB/scalars"
	graphql "github.com/neelance/graphql-go"
)

type Connection struct {
	First *int32
	After *scalars.Cursor
}

type ID struct {
	ID *graphql.ID
}

type Search struct {
	ID   *graphql.ID
	Name *string
}

type FlavorText struct {
	VersionGroup int32
}
