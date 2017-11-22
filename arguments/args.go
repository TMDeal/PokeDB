package arguments

import (
	"github.com/TMDeal/PokeDB/scalars"
)

type Connection struct {
	First *int32
	After *scalars.Cursor
}

type Name struct {
	Name *string
}

type ID struct {
	ID *int32
}

type Search struct {
	Name
	ID
}
