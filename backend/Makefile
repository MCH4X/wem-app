.PHONY: lint test run build dev clean migration-up migration-down seed

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

migration-up:
	go run ./cmd/migrate up

migration-down:
	go run ./cmd/migrate down

seed:
	go run ./cmd/seed