.PHONY: lint test run dev

lint:
	golangci-lint run -v

test:
	go test ./...

run:
	go run ./cmd

dev: lint test
	@echo "✅ All checks passed"
	@printf "\a"