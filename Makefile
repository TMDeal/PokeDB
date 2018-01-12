SERVER_OUT := pokedb
CONNECTION_OUT := connection
GRAPHQL_OUT := schema/assets.go
GO_BINDATA_BINARY := $(GOPATH)/bin/go-bindata
GOOSE_BINARY := $(GOPATH)/bin/goose
GOOSE_DIR := "./migrations"
GOOSE_DB_DRIVER := postgres
GOOSE_DB_INFO := "user=pokedb dbname=pokedb sslmode=disable"

.PHONY: all
all: pokedb

pokedb: vendor generate
	@go build -o ${SERVER_OUT} cmd/pokedb/main.go
	@echo "built $(SERVER_OUT) binary"

connection:
	@go build -o ${CONNECTION_OUT} cmd/connection/main.go
	@echo "built $(CONNECTION_OUT) binary"

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

$(GOOSE_BINARY):
	@go get -u github.com/pressly/goose/cmd/goose
	@echo "Installed goose dependency"

.PHONY: migrate
migrate-up: $(GOOSE_BINARY)
	@goose -dir=$(GOOSE_DIR) $(GOOSE_DB_DRIVER) $(GOOSE_DB_INFO) up
	@echo "Migration up complete"

migrate-down: $(GOOSE_BINARY)
	@goose -dir=$(GOOSE_DIR) $(GOOSE_DB_DRIVER) $(GOOSE_DB_INFO) down-to 00001
	@echo "Migration down complete"

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
