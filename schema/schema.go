package schema

import (
	"bytes"
	"log"
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go graphql/

func MustRead() string {
	schema, err := Read()
	if err != nil {
		log.Fatal(err)
	}

	return schema
}

func Read() (string, error) {
	var schema []byte

	for _, f := range AssetNames() {
		data, err := Asset(f)
		if err != nil {
			return "", err
		}

		schema = bytes.Join(
			[][]byte{
				schema,
				data,
			},
			[]byte{},
		)
	}

	return string(schema), nil
}
