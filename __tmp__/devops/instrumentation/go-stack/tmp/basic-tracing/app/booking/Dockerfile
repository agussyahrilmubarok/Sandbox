FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY build/booking-app /app/booking-app
COPY config-prod.json /app/configs/config.json

EXPOSE 8080

ENTRYPOINT ["/app/booking-app"]
