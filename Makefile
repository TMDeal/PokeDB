SERVER_OUT := pokedb
CONNECTION_OUT := connection
GRAPHQL_OUT := schema/assets.go
GO_BINDATA_BINARY := $(GOPATH)/bin/go-bindata

.PHONY: all
all: pokedb

pokedb: generate
	go build -o ${SERVER_OUT} cmd/pokedb/main.go

connection:
	go build -o ${CONNECTION_OUT} cmd/connection/main.go

.PHONY: run
run: generate
	go run ./cmd/pokedb/main.go

.PHONY: test
test:
	go test ./...

$(GO_BINDATA_BINARY):
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: generate
generate: connection $(GO_BINDATA_BINARY)
	go generate ./resolvers
	go generate ./schema

.PHONY: clean
clean:
	rm ${SERVER_OUT}
	rm ${CONNECTION_OUT}
	rm ${GRAPHQL_OUT}
	rm ./resolvers/*_connection.go
