package schema

const schema = `
schema {
    query: Query
}

type Query {
    #Searches
    generation(id: Int!): Generation
    # region(id: Int, name: String): Region
    # type(id: Int, name: String): Type

    # Lists
    generations(first: Int = 20): [Generation]
    regions(first: Int = 20): [Region]
    types(first: Int = 20): [Type]
}
`
