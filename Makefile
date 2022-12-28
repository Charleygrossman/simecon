.SILENT:

test:
	go test ./...

build:
	make build-gen && make build-sim

build-gen:
	go build -o bin/gen ./cmd/gen

build-sim:
	go build -o bin/sim ./cmd/sim