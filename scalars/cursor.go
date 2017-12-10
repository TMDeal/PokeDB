package scalars

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Cursor string

func NewCursor(t string, i int) Cursor {
	v := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s-%d", t, i)))
	return Cursor(v)
}

func (c Cursor) IntValue() (int, error) {
	b, err := base64.StdEncoding.DecodeString(string(c))
	if err != nil {
		return 0, err
	}

	elems := strings.Split(string(b), "-")

	i, err := strconv.Atoi(elems[1])
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

		elems := strings.Split(string(b), "-")

		i, err := strconv.Atoi(elems[1])
		if err != nil {
			return err
		}

		*c = NewCursor(elems[0], i)
	default:
		err = errors.New("wrong type")
	}

	return err
}
