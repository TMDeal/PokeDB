package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/scalars"
)

//CResolver is an interface that says how to get the total count of items in a
//page query, and the pageinfo of a query for a connection
type CResolver interface {
	//TotalCount returns the number of items in a page
	TotalCount() (int32, error)
	//Pageinfo returns the information about a page regarding its starting and
	//ending items, and knows if there is a next page or not
	PageInfo() (*PageResolver, error)
}

//EResolver is an interface that says that an Edge must have a Cursor item
type EResolver interface {
	//Cursor returns the cursor for the edge
	Cursor() scalars.Cursor
}

//MakeCursors creates two cursors, start and end, that corrospond to the page
//queries first element and last element.
func MakeCursors(size int, args arguments.Connection) (*scalars.Cursor, *scalars.Cursor, error) {
	start := scalars.NewCursor(0)
	if args.After != nil {
		start = *args.After
	}

	starti, err := start.IntValue()
	if err != nil {
		return nil, nil, err
	}

	end := scalars.NewCursor(size + starti)
	if args.First != nil {
		actualEnd := starti + int(*args.First)
		if actualEnd > size {
			end = scalars.NewCursor(actualEnd)
		}
	}

	return &start, &end, nil
}

//HasNextPage is a helper function to tell if there is a next page in the
//pagination query.
func HasNextPage(end scalars.Cursor, size int) (bool, error) {
	endi, err := end.IntValue()
	if err != nil {
		return false, err
	}

	hasNext := endi < size
	return hasNext, nil
}
