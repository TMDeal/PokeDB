package resolvers

import graphql "github.com/neelance/graphql-go"

type Node interface {
	ID() graphql.ID
}
