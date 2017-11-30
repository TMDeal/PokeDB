package schema

const schema = `
schema {
    query: Query
}

scalar Cursor

type Query {
    #Searches
	node(id: ID): Node
    region(id: ID, name: String): Region
	#generation(id: ID): Generation
    # type(id: ID, name: String): Type

    # Lists
	regions(first: Int, after: Cursor): RegionConnection!
	generations(first: Int, after: Cursor): GenerationConnection!
	types(first: Int, after: Cursor): TypeConnection!
}
`
