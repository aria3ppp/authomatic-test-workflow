ENV_FILE ?= .env
POSTGRES_USER ?= $(shell sed -r -n 's/POSTGRES_USER="(.+)"/\1/p' $(ENV_FILE))
POSTGRES_PASSWORD ?= $(shell sed -r -n 's/POSTGRES_PASSWORD="(.+)"/\1/p' $(ENV_FILE))
POSTGRES_PORT ?= $(shell sed -r -n 's/POSTGRES_PORT="(.+)"/\1/p' $(ENV_FILE))
POSTGRES_DB ?= $(shell sed -r -n 's/POSTGRES_DB="(.+)"/\1/p' $(ENV_FILE))
APP_DSN ?= "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --user "$(shell id -u):$(shell id -g)" --network host migrate/migrate:4 -path=/migrations -database "$(APP_DSN)"

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: db-start
db-start: ## start the database server
	docker run --rm --name postgres \
		-e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -e POSTGRES_PORT=$(POSTGRES_PORT) \
		-d -p 5432:5432 postgres:14-alpine

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop postgres

.PHONY: elk-start
elk-start: ## start elk
	docker run --rm --name elk \
		-e discovery.type=single-node -e xpack.security.enabled=false -e bootstrap.memory_lock=true \
		-d -p 9200:9200 elasticsearch:8.4.0

.PHONY: elk-stop
elk-stop: ## stop elk
	docker stop elk

.PHONY: test-all
test-all: ## run all tests
	@echo "mode: count" > coverage-all.out
	make test-unit
	@tail -n +2 coverage.out >> coverage-all.out
	make test-models
	@tail -n +2 coverage.out >> coverage-all.out
	make test-integration
	@tail -n +2 coverage.out >> coverage-all.out
	make test-e2e
	@tail -n +2 coverage.out >> coverage-all.out
	# env TEST_DB_INTEGRATION=V TEST_E2E=V go test -covermode=count -coverprofile=coverage.out `go list ./... | grep -v /models`

.PHONY: test-unit
test-unit: ## run unit tests
	go test -covermode=count -coverprofile=coverage.out `go list ./... | grep -v /models`

.PHONY: test-models
test-models: ## run models tests
	make migrate
	go test -covermode=count -coverprofile=coverage.out  ./internal/models/...

.PHONY: test-integration
test-integration: ## run integration tests
	env TEST_DB_INTEGRATION=V go test -covermode=count -coverprofile=coverage.out ./internal/repo/

.PHONY: test-e2e
test-e2e: ## run end-to-end tests
	env TEST_E2E=V go test -covermode=count -coverprofile=coverage.out ./internal/server/

.PHONY: test-all-cover
test-all-cover: test-all ## run all tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: test-unit-cover
test-unit-cover: test-unit ## run unit tests and show test coverage information
	go tool cover -html=coverage.out

.PHONY: test-models-cover
test-models-cover: test-models ## run models tests and show test coverage information
	go tool cover -html=coverage.out

.PHONY: test-integration-cover
test-integration-cover: test-integration ## run integration tests and show test coverage information
	go tool cover -html=coverage.out

.PHONY: test-e2e-cover
test-e2e-cover: test-e2e ## run end-to-end tests and show test coverage information
	go tool cover -html=coverage.out

.PHONY: test-arg-cover
test-arg-cover: ## run tests by passing $ARG env value to 'go test' command and show test coverge information
	go test -covermode=count -coverprofile=coverage.out $(ARG)
	go tool cover -html=coverage.out

.PHONY: run
run: ## run main package
	go run .

.PHONY: build
build: ## build main package
	go build .

.PHONY: generate
generate: ## run 'go generate' for all packages
	go generate ./...

.PHONY: lint
lint: ## run staticcheck
	@staticcheck ./...

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migrate step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $${name}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop -f
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-arg
migrate-arg: ## run migration command with argument ARG
	@echo "Running migration command with argument: $(ARG)"
	@$(MIGRATE) $(ARG)
