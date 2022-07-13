.SILENT:

build:
	go mod download && CGO_ENABLED GODS=linux go build -o ./bin/app ./cmd/app/main.go
run: build
