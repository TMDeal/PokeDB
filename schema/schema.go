package schema

import "strings"

func New() string {
	schema := strings.Join(
		[]string{
			schema,
			interfaces,
			pages,
			generations,
			regions,
			types,
			stats,
			moves,
			contests,
		},
		"",
	)

	return schema
}
