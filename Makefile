SERVER_OUT := pokedb
CONNECTION_OUT := connection
GRAPHQL_OUT := schema/assets.go
GO_BINDATA_BINARY := $(GOPATH)/bin/go-bindata

.PHONY: all
all: pokedb

pokedb: vendor generate
	@go build -o ${SERVER_OUT} cmd/pokedb/main.go
	@echo "built pokedb binary"

connection:
	@go build -o ${CONNECTION_OUT} cmd/connection/main.go
	@echo "built connection binary"

vendor:
	@glide install

.PHONY: run
run: generate
	@go run ./cmd/pokedb/main.go

.PHONY: test
test:
	@go test ./...

$(GO_BINDATA_BINARY):
	@go get -u github.com/jteeuwen/go-bindata/...
	@echo "Installed go-bindata dependency"

.PHONY: generate
generate: connection $(GO_BINDATA_BINARY)
	@go generate ./resolvers
	@echo "Generated resolvers"
	@go generate ./schema
	@echo "Generated schema"

.PHONY: clean
clean:
	@rm -f ${SERVER_OUT}
	@rm -f ${CONNECTION_OUT}
	@rm -f ${GRAPHQL_OUT}
	@rm -f ./resolvers/*_connection.go
