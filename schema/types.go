package schema

const types = `
type Type implements Node {
	id: ID!
	identifier: String!
	name: String!
	generation: Generation!
	damageClass: DamageClass
	efficacy: TypeEfficacy!
}

type TypeEfficacyDirection {
	from: [Type]!
	to: [Type]!
}

type TypeEfficacy {
	double: TypeEfficacyDirection!
	normal: TypeEfficacyDirection!
	half: TypeEfficacyDirection!
	none: TypeEfficacyDirection!
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
