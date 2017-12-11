package schema

const moves = `
type Move implements Node {
	id: ID!
	identifier: String!
	name: String!
	generation: Generation!
	damageClass: DamageClass!
	contestEffect: ContestEffect
	contestType: ContestType
	superContestEffect: SuperContestEffect
	flavorText(versionGroup: Int = 17): String!
	effect: MoveEffect
	target: MoveTarget!
	flags: [MoveFlag]
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

type MoveEffect {
	short: String!
	long: String!
}

type MoveTarget {
    identifier: String!
    name: String!
    description: String!
}

type MoveFlag {
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
