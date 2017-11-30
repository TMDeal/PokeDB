package schema

const types = `
type Type implements Node {
    id: ID!
    identifier: String!
    name: String!
    generation: Generation!
    damageClass: DamageClass
}

type TypeEdge {
	cursor: Cursor!
	node: Type
}

type TypeConnection {
	totalCount: Int!
	edges: [TypeEdge]
	pageInfo: PageInfo!
}
`
