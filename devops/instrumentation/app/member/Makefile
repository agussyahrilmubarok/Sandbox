server/dev:
	go run cmd/server/main.go --config=config.json

server/build:
	rm -rf build/
	go build -o build/member-app cmd/server/main.go

swagger/install:
	go install github.com/swaggo/swag/cmd/swag@latest

swagger/generate:
	swag init -g ./cmd/server/main.go -o ./cmd/server/docs