app/dev:
	air

app/start:
	go run cmd/finly/main.go

app/seed:
	MONGO_URI=mongodb://admin:securepassword123@localhost:27017 go run cmd/seed/main.go

docker/compose/dev/up:
	docker compose -f docker/docker-compose-dev.yml up -d

docker/compose/dev/down:
	docker compose -f docker/docker-compose-dev.yml down
