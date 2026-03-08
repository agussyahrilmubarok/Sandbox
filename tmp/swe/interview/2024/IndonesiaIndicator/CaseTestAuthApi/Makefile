dev:
	./mvnw spring-boot:run -DskipTests

test:
	./mvnw test

test.coverage:
	./mvnw clean test jacoco:report

db.dev:
	docker compose -f docker-mysql.yml up