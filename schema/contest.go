package schema

const contests = `
type ContestEffect {
	jam: Int!
	appeal: Int!
	flavorText: String!
	effect: String!
}

type ContestType {
	identifier: String!
	name: String!
	flavor: String!
	color: String!
}

type SuperContestEffect {
	appeal: String!
	flavorText: String!
}
`
