app.run:
	go run cmd/app/main.go

app.test:
	go test ./internal/... -coverprofile=docs/tests/coverage.out

app.build:
	go build -o bin/app cmd/app/main.go

docker.up:
	docker compose up -d

docker.down:
	docker compose down

migrater.up:
	go run cmd/migrater/main.go up

migrater.down:
	go run cmd/migrater/main.go down

seeder.up:
	go run cmd/seeder/main.go

seeder.down:
	go run cmd/seeder/main.go down

migrater.build:
	go build -o bin/migrater cmd/migrater/main.go