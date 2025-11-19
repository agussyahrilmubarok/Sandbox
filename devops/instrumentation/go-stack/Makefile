app/all/build:
	cd app/member && \
	go fmt ./... && \
	go vet ./... && \
	go build -o build/member-app cmd/server/main.go
	cd app/course && \
	go fmt ./... && \
	go vet ./... && \
	go build -o build/course-app cmd/server/main.go
	cd app/booking && \
	go fmt ./... && \
	go vet ./... && \
	go build -o build/booking-app cmd/server/main.go

compose/up:
	docker compose up -d --build

compose/down:
	docker compose down
	docker volume prune -a

compose/dev/up:
	docker compose -f docker-compose-dev.yml up -d --build

compose/dev/down:
	docker compose -f docker-compose-dev.yml down
	docker volume prune -a

docker/image/rm/norepo:
	docker images -q --filter "dangling=true" | xargs docker rmi
