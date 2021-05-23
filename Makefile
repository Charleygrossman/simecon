.SILENT:

build:
	go build -o bin/tradesim

run:
	go run main.go

test:
	go test ./...