# Load environment variables from .env
include .env
export

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: run the api server
.PHONY: run
run:
	@go run cmd/api/main.go

## run/live: run the api server live via air
.PHONY: run/live
run/live:
	@air

## build: build the api server
.PHONY: build
build:
	@CGO_ENABLED=0 go build -o tmp/app cmd/api/main.go

## swag: render swagger on port 3000
.PHONY: swag
swag:
	@sh -c "cd api/swagger-ui && npm start"

## test: run all tests
.PHONY: test
test:
	@sh -c "cd test && go test -v -race -buildvcs"

# ==================================================================================== #
# STAGING & PRODUCTION
# ==================================================================================== #

## fetchchanges: fetch changes
.PHONY: fetchchanges
fetchchanges:
	@git checkout master
	@git fetch origin
	@git merge origin/master

## stage: stage api
.PHONY: stage
stage:
	@docker compose --file compose.yaml --file compose.stage.yaml up --build backend --no-deps --detach backend

## freshdb: freshdb
.PHONY: freshdb
freshdb:
	@docker compose --file compose.yaml down db
	@docker volume rm avanzu_db-persistence
	@docker compose --file compose.yaml up --no-deps --detach db

# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

DB_DSN="mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_SCHEMA}"

## mig/new name=$1: create a new database migration
.PHONY: mig/new
mig/new:
ifdef table
	@migrate create -dir db/migrations -ext sql ${table}
else
	@echo "must define \`table\` argument"
endif

## mig/up: apply all up database migrations
.PHONY: mig/up
mig/up:
	@docker exec -it db mariadb -u root -e 'DROP SCHEMA locavest; CREATE SCHEMA locavest;'
	@migrate -path=db/migrations -database=${DB_DSN} up

## mig/down: apply all down database migrations
.PHONY: mig/down
mig/down:
	@migrate -path=db/migrations -database=${DB_DSN} down -all

## mig/goto version=$1: migrate to a specific version number
.PHONY: mig/goto
mig/goto:
	@migrate -path=db/migrations -database=${DB_DSN} goto ${version}

## mig/f version=$1: force database migration
.PHONY: mig/f
mig/f:
	@migrate -path=db/migrations -database=${DB_DSN} force ${version}

## mig/v: print the current in-use migration version
.PHONY: mig/v
mig/v:
	@migrate -path=db/migrations -database=${DB_DSN} version

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
