OUT := pokedb

build: generate
	go build -o ${OUT} cmd/server/pokedb.go

run: build
	./${OUT}

test: build
	go test ./...

generate:
	go generate ./resolvers

clean:
	rm ${OUT}
	rm ./resolvers/*_connection.go

.PHONY: run build test generate