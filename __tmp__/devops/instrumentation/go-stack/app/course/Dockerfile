FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY build/course-app /app/course-app
COPY config-prod.json /app/configs/config.json

EXPOSE 8082

ENTRYPOINT ["/app/course-app"]
