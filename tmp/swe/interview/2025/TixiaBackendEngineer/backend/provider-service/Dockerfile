FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o provider-service ./cmd/main.go

FROM alpine:3.20
RUN apk add --no-cache tzdata ca-certificates
WORKDIR /root/
COPY --from=builder /app/provider-service .
COPY --from=builder /app/tmp ./tmp

CMD ["./provider-service", "--serviceLevel=production"]