package schema

import "strings"

func New() string {
	schema := strings.Join(
		[]string{
			schema,
			pages,
			generations,
			regions,
			types,
			moves,
		},
		"",
	)

	return schema
}
