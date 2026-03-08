FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY build/member-app /app/member-app
COPY config-prod.json /app/configs/config.json

EXPOSE 8081

ENTRYPOINT ["/app/member-app"]
