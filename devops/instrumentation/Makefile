compose/up:
	docker compose up -d --build

compose/down:
	docker compose down
	docker volume prune -a

docker/image/rm/norepo:
	docker images -q --filter "dangling=true" | xargs docker rmi
