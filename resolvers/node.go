package resolvers

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	graphql "github.com/neelance/graphql-go"
)

type Node interface {
	ID() graphql.ID
}

type NodeResolver struct {
	Node
}

func NewNodeResolver(node Node) *NodeResolver {
	return &NodeResolver{node}
}

func (r NodeResolver) ToRegion() (*RegionResolver, bool) {
	to, ok := r.Node.(*RegionResolver)
	return to, ok
}

func (r NodeResolver) ToMove() (*MoveResolver, bool) {
	to, ok := r.Node.(*MoveResolver)
	return to, ok
}

func (r *NodeResolver) ToGeneration() (*GenerationResolver, bool) {
	to, ok := r.Node.(*GenerationResolver)
	return to, ok
}

func (r *NodeResolver) ToStat() (*StatResolver, bool) {
	to, ok := r.Node.(*StatResolver)
	return to, ok
}

func (r *NodeResolver) ToType() (*TypeResolver, bool) {
	to, ok := r.Node.(*TypeResolver)
	return to, ok
}

func (r *NodeResolver) ToVersion() (*VersionResolver, bool) {
	to, ok := r.Node.(*VersionResolver)
	return to, ok
}

func (r *NodeResolver) ToVersionGroup() (*VersionGroupResolver, bool) {
	to, ok := r.Node.(*VersionGroupResolver)
	return to, ok
}

func (r *NodeResolver) ToAbility() (*AbilityResolver, bool) {
	to, ok := r.Node.(*AbilityResolver)
	return to, ok
}

func (r *NodeResolver) ToItem() (*ItemResolver, bool) {
	to, ok := r.Node.(*ItemResolver)
	return to, ok
}

func GlobalID(t interface{}, id int64) graphql.ID {
	value := reflect.ValueOf(t)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	base := []byte(fmt.Sprintf("%s-%d", value.Type().Name(), id))
	encrypted := base64.StdEncoding.EncodeToString(base)

	return graphql.ID(encrypted)
}

func FromGlobalID(id graphql.ID) (string, int, error) {
	b, err := base64.StdEncoding.DecodeString(string(id))
	if err != nil {
		return "", 0, err
	}

	elems := strings.Split(string(b), "-")

	i, err := strconv.Atoi(elems[1])
	if err != nil {
		return "", 0, err
	}

	return elems[0], i, nil
}
