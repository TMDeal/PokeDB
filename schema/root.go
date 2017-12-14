package schema

const schema = `
schema {
	query: Query
}

scalar Cursor

type Query {
	node(id: ID): Node

	ability(id: ID, name: String): Ability
	region(id: ID, name: String): Region
	generation(id: ID, name: String): Generation
	type(id: ID, name: String): Type
	move(id: ID, name: String): Move
	version(id: ID, name: String): Version
	versionGroup(id: ID, name: String): VersionGroup

	abilities(first: Int, after: Cursor): AbilityConnection!
	regions(first: Int, after: Cursor): RegionConnection!
	generations(first: Int, after: Cursor): GenerationConnection!
	types(first: Int, after: Cursor): TypeConnection!
	moves(first: Int, after: Cursor): MoveConnection!
	versions(first: Int, after: Cursor): VersionConnection!
	versionGroups(first: Int, after: Cursor): VersionGroupConnection!
}
`
