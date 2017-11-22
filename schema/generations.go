package schema

const generations = `
type Generation {
    id: ID!
    identifier: String!
    name: String!
    region: Region!
}

type GenerationEdge {
	cursor: Cursor!
	node: Region
}

type GenerationConnection {
	totalCount: Int!
	edges: [GenerationEdge]
	pageInfo: PageInfo!
}
`
