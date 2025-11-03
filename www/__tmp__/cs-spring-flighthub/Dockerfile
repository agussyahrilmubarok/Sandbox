FROM maven:3.9.7-amazoncorretto-21 AS build
COPY pom.xml ./
COPY .mvn .mvn
COPY src src
RUN mvn clean install -DskipTests

FROM amazoncorretto:21
WORKDIR flighthub
COPY --from=build target/*.jar flighthub.jar
EXPOSE 1133

ENTRYPOINT ["java", "-jar", "flighthub.jar"]