package models

import "errors"

//ErrInvalidSearch is returned when a query is given an invalid search argument
var ErrInvalidSearch = errors.New("Invalid search for query. Must supply int or string")
