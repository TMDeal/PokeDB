package schema

const schema = `
schema {
	query: Query
}

scalar Cursor

type Query {
	# Searches
	node(id: ID): Node
	region(id: ID, name: String): Region
	#generation(id: ID, name: String) Generation
	#type(id: ID, name: String) Type
	move(id: ID, name: String): Move

	# Connections
	regions(first: Int, after: Cursor): RegionConnection!
	generations(first: Int, after: Cursor): GenerationConnection!
	types(first: Int, after: Cursor): TypeConnection!
	moves(first: Int, after: Cursor): MoveConnection!
}
`
