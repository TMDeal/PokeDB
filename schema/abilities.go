package schema

var abilities = `
type Ability implements Node {
	id: ID!
	identifier: String!
	name: String!
	shortEffect: String!
	effect: String!
	generation: Generation!
	flavorText(versionGroup: Int = 17): AbilityFlavorText!
}

type AbilityEdge {
	cursor: Cursor!
	node: Ability
}

type AbilityConnection {
	totalCount: Int!
	pageInfo: PageInfo!
	edges: [AbilityEdge]
}

type AbilityFlavorText implements FlavorText {
	text: String!
	versionGroup: VersionGroup
}
`
