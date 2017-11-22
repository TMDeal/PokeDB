package types

import (
	"github.com/TMDeal/PokeDB/scalars"
)

type EdgeResolver struct {
	node   *Resolver
	cursor scalars.Cursor
}

func NewEdgeResolver(n *Resolver, c scalars.Cursor) *EdgeResolver {
	return &EdgeResolver{
		node:   n,
		cursor: c,
	}
}

func (e *EdgeResolver) Cursor() scalars.Cursor {
	return e.cursor
}

func (e *EdgeResolver) Node() *Resolver {
	return e.node
}
