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
	flavorText(versionGroup: Int = 17): MoveFlavortext!
	effect: MoveEffect
	target: MoveTarget!
	flags: [MoveFlag]
	meta: MoveMeta!
	type: Type!
	power: Int
	pp: Int
	accuracy: Int
	priority: Int!
}

type MoveMeta {
	minHits: Int
	maxHits: Int
	minTurns: Int
	maxTurns: Int
	drain: Int!
	healing: Int!
	critRate: Int!
	ailmentChance: Int!
	flinchChance: Int!
	statChance: Int!
	ailment: MoveMetaAilment!
	category: MoveMetaCategory!
	statChanges: [MoveMetaStatChange]!
}

type MoveMetaAilment {
	identifier: String!
	name: String!
}

type MoveMetaCategory {
	identifier: String!
	description: String!
}

type MoveMetaStatChange {
	change: Int!
	stat: Stat!
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

type MoveFlavortext implements FlavorText {
	text: String!
	versionGroup: VersionGroup
}

type DamageClass {
    identifier: String!
    name: String!
    description: String!
}
`
