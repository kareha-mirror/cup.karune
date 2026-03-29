all: build

build:
	go build -o karune ./cmd/karune

clean:
	rm -f karune

run:
	go run ./cmd/karune

fmt:
	go fmt ./...

test:
	go test ./...
