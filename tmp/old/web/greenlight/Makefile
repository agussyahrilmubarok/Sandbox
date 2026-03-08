# ==================================================================================== #
# VARIABLES
# ==================================================================================== #

# Environment Variables
GREENLIGHT_DB_DSN ?= postgres://greenlight:my-secret-pw@localhost/greenlight?sslmode=disable
DB_PW ?= my-secret-pw

# Production Server Information
production_host_ip = 144.126.210.226

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/  /'

.PHONY: confirm
confirm:
	@echo 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@echo 'Starting the API application...'
	@go run ./cmd/api -db-dsn=${GREENLIGHT_DB_DSN}

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	@echo 'Connecting to the database...'
	@psql ${GREENLIGHT_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@test -n "$(name)" || (echo "ERROR: Migration name not provided"; exit 1)
	@echo 'Creating migration files for ${name}'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	@migrate -path="./migrations" -database "${GREENLIGHT_DB_DSN}" up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet, and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	@go fmt ./...
	@echo 'Vetting code...'
	@go vet ./...
	@echo 'Running static analysis...'
	@staticcheck ./...
	@echo 'Running tests...'
	@go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	@go mod tidy
	@go mod verify
	@echo 'Vendoring dependencies...'
	@go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api for local OS...'
	@go build -ldflags '-s -w' -o ./bin/api ./cmd/api
	@echo 'Building cmd/api for Linux amd64...'
	@GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o ./bin/linux_amd64/api ./cmd/api

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	@echo 'Connecting to production server...'
	@ssh greenlight@${production_host_ip}

## production/deploy/api: deploy the API to production
.PHONY: production/deploy/api
production/deploy/api:
	@echo 'Deploying API to production...'
	@rsync -P ./bin/linux_amd64/api greenlight@${production_host_ip}:~
	@rsync -rP --delete ./migrations greenlight@${production_host_ip}:~
	@rsync -P ./remote/production/api.service greenlight@${production_host_ip}:~
	@rsync -P ./remote/production/Caddyfile greenlight@${production_host_ip}:~
	@ssh -t greenlight@${production_host_ip} '\
		migrate -path ~/migrations -database $$GREENLIGHT_DB_DSN up \
        && sudo mv ~/api.service /etc/systemd/system/ \
        && sudo systemctl enable api \
        && sudo systemctl restart api \
        && sudo mv ~/Caddyfile /etc/caddy/ \
        && sudo systemctl reload caddy \
	'

# ==================================================================================== #
# DOCKER
# ==================================================================================== #

## docker.dev.up: start the development Docker environment
.PHONY: docker.dev.up
docker.dev.up:
	@echo 'Starting Docker development environment...'
	@docker compose -f docker-compose.dev.yml up

## docker.dev.down: stop the development Docker environment
.PHONY: docker.dev.down
docker.dev.down:
	@echo 'Stopping Docker development environment...'
	@docker compose -f docker-compose.dev.yml down
