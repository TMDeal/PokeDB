SERVER_OUT := pokedb
GENERATOR_OUT := connection

.PHONY: all
all: pokedb

pokedb: generate
	go build -o ${SERVER_OUT} cmd/pokedb/main.go

connection:
	go build -o ${GENERATOR_OUT} cmd/connection/main.go

.PHONY: run
run: generate
	go run ./cmd/pokedb/main.go

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate: connection
	go generate ./resolvers

.PHONY: clean
clean:
	rm ${OUT}
	rm ./resolvers/*_connection.go
