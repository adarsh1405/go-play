## Makefile for a small Go project
## Usage: make [target]

.PHONY: all help build run test fmt vet tidy clean

all: build

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build   - build the binary to ./bin/app"
	@echo "  run     - run the app"
	@echo "  test    - run go tests"
	@echo "  fmt     - go fmt ./..."
	@echo "  vet     - go vet ./..."
	@echo "  tidy    - go mod tidy"
	@echo "  clean   - remove bin/"

build:
	@mkdir -p bin
	go build -v -o bin/app ./...

run:
	go run ./...

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf bin
