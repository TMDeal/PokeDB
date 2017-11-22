package schema

const regions = `
type Region {
    id: ID!
    identifier: String!
    name: String!
}

type RegionEdge {
	cursor: Cursor!
	node: Region
}

type RegionConnection {
	totalCount: Int!
	edges: [RegionEdge]
	pageInfo: PageInfo!
}
`
