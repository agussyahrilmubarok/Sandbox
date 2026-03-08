################################################################################
##### DOCKER
################################################################################
docker/compose/up:
	docker compose up -d

docker/compose/down:
	docker compose down

docker/image/build:
	docker build -t finly .

docker/prune:
	docker container prune && docker volume prune && docker image prune