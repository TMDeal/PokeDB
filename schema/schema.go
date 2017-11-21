package schema

import "strings"

func New() string {
	schema := strings.Join(
		[]string{
			schema,
			generations,
			regions,
			types,
			moves,
		},
		"",
	)

	return schema
}
