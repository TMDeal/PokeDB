package schema

const interfaces = `
interface Node {
	id: ID!
}

interface FlavorText {
	text: String!
	versionGroup: VersionGroup
}
`
