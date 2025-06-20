.PHONY: all run clean fmt build

all: run


run:
	@echo "Running the project..."
	go run ./cmd/main.go

build:
	@echo "Building the project..."
	go build -o bin/main ./cmd

clean:
	@echo "Cleaning up..."
	rm -rf bin

fmt:
	@echo "Formatting code..."
	go fmt ./...