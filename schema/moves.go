package schema

const moves = `
type Move implements Node {
	id: ID!
	identifier: String!
	name: String!
	generation: Generation!
	damageClass: DamageClass!
	type: Type!
	power: Int
	pp: Int
	accuracy: Int
	priority: Int!
}

type MoveEdge {
	cursor: Cursor!
	node: Move
}

type MoveConnection {
	totalCount: Int!
	edges: [MoveEdge]
	pageInfo: PageInfo!
}

type MoveTarget {
    identifier: String!
    name: String!
    description: String!
}

type DamageClass {
    identifier: String!
    name: String!
    description: String!
}
`
