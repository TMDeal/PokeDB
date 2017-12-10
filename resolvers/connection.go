package resolvers

import (
	"github.com/TMDeal/PokeDB/arguments"
	"github.com/TMDeal/PokeDB/scalars"
)

//MakeCursors creates two cursors, start and end, that corrospond to the page
//queries first element and last element.
func MakeCursors(t string, size int, args arguments.Connection) (*scalars.Cursor, *scalars.Cursor, error) {
	start := scalars.NewCursor(t, 0)
	if args.After != nil {
		start = *args.After
	}

	starti, err := start.IntValue()
	if err != nil {
		return nil, nil, err
	}

	end := scalars.NewCursor(t, size+starti)
	if args.First != nil {
		actualEnd := starti + int(*args.First)
		if actualEnd > size {
			end = scalars.NewCursor(t, actualEnd)
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
