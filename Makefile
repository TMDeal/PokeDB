OUT := pokedb

build: generate
	go build -o ${OUT} cmd/pokedb/main.go

run: build
	./${OUT}

test: build
	go test ./...

generate:
	go generate ./resolvers

clean:
	rm ${OUT}
	rm ./resolvers/*_connection.go

.PHONY: build run test generate clean
