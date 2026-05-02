.PHONY: lint test run build dev clean

lint:
	golangci-lint run -v

test:
	go test ./...

run:
	go run ./cmd

build:
	go build -o bin/api ./cmd

clean:
	rm -rf bin/

dev: lint test build
	@echo "✅ All checks passed"