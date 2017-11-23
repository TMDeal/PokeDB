package scalars

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Cursor string

func NewCursor(i int) Cursor {
	v := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i)))
	return Cursor(v)
}

func (c Cursor) IntValue() (int, error) {
	b, err := base64.StdEncoding.DecodeString(string(c))
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (_ Cursor) ImplementsGraphQLType(name string) bool {
	return name == "Cursor"
}

func (c *Cursor) UnmarshalGraphQL(input interface{}) error {
	var err error

	switch input := input.(type) {
	case string:
		b, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			return err
		}

		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return err
		}

		*c = NewCursor(i)
	case int32:
		*c = NewCursor(int(input))
	default:
		err = errors.New("wrong type")
	}

	return err
}
