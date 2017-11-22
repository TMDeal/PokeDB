package schema

const schema = `
schema {
    query: Query
}

scalar Cursor

type Query {
    #Searches
    generation(id: Int): Generation
    # region(id: Int, name: String): Region
    # type(id: Int, name: String): Type

    # Lists
	regions(first: Int, after: Cursor): RegionConnection!
	generations(first: Int, after: Cursor): GenerationConnection!
	types(first: Int, after: Cursor): TypeConnection!
}
`
