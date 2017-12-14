package schema

var versions = `
type Version implements Node {
	id: ID!
	identifier: String!
	name: String!
	group: VersionGroup!
}

type VersionEdge {
	cursor: Cursor!
	node: Version!
}

type VersionConnection {
	totalCount: Int!
	pageInfo: PageInfo!
	edges: [VersionEdge]
}

type VersionGroup implements Node {
	id: ID!
	identifier: String!
	ordering: Int!
	generation: Generation!
	versions: [Version]!
	regions: [Region]!
}

type VersionGroupEdge {
	cursor: Cursor!
	node: VersionGroup!
}

type VersionGroupConnection {
	totalCount: Int!
	pageInfo: PageInfo!
	edges: [VersionGroupEdge]
}
`
