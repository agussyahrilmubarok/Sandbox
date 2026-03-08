app.dev:
	mvn clean spring-boot:run

app.package:
	mvn clean install package

app.image:
	docker build -t mzs-house/flighthub .

kube.start:
	minikube start

kube.dashboard:
	minikube dashboard

kube.ingress:
	minikube addons enable ingress

kube.ingress.check:
	kubectl get svc -n ingress-nginx

kube.apply:
	eval $(minikube docker-env)
	docker build -t mzs-house/flighthub .
	kubectl apply -f ./scripts/k8s

kube.delete:
	kubectl delete -f ./scripts/k8s

kube.prome.service:
	minikube service prometheus-service

kube.grafana.service:
	minikube service grafana-service

kube.connect.image.docker:
	eval $(minikube docker-env)

kube.node.forward:
	kubectl port-forward service/flighthub-service 1133:1133
	kubectl port-forward service/grafana-service 3000:3000
	kubectl port-forward service/prometheus-service 9090:9090

docker.dev.up:
	docker compose -f docker-compose-dev.yml up -d --build

docker.dev.down:
	docker compose -f docker-compose-dev.yml down
	docker volume rm cs-spring-flighthub_mongodb_data_container

docker.prod.up:
	docker compose up -d --build

docker.prod.down:
	docker compose down
	docker volume rm cs-spring-flighthub_mongodb_data_container
	docker volume rm cs-spring-flighthub_data
	docker volume rm cs-spring-prometheus_data
	docker volume rm cs-spring-grafana_data

mongo.bash:
	docker exec -it mongo-container bash

mongo.connect:
	mongosh -u <username> -p <password> --authenticationDatabase admin

mongo.db.show:
	show dbs

mongo.collection.show:
	show collections

mongo.select:
	db.users.find().pretty()