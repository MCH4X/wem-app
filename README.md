# wem-app

WEM — Where's My Money? 💸 Track it before it's gone again. 🔍

## Project Structure

```
wem-app/
├── backend/          # Go API server
│   ├── cmd/          # Application entry point
│   ├── internal/     # Private application code
│   │   ├── handlers/ # HTTP handlers
│   │   └── routes/   # Route definitions
│   ├── Makefile      # Build & dev commands
│   └── .golangci.yml # Linter configuration
└── .github/
    └── workflows/
        └── backend.yml  # CI pipeline
```

## Backend

### Prerequisites

- [Go 1.26+](https://go.dev/dl/)
- [golangci-lint](https://golangci-lint.run/welcome/install/)

### Getting Started

```bash
cd backend

# Run the server
make run

# Run linter
make lint

# Run tests
make test

# Build binary
make build

# Run all checks (lint → test → build)
make dev

# Clean build artifacts
make clean
```

### CI Pipeline

The backend CI runs automatically on every push and pull request to `main`.  
It uses [`dorny/paths-filter`](https://github.com/dorny/paths-filter) to detect changes — CI steps only run when `backend/` files are modified, but a status is **always returned** to avoid Branch Protection deadlocks.

```
Every PR
  │
  ▼
changes (detect modified paths)
  │
  ├── backend changed → Lint → Test → Build
  └── no backend changes → skip (success)
```

| Step  | Tool              | Description                    |
|-------|-------------------|--------------------------------|
| Lint  | golangci-lint     | Code quality & static analysis |
| Test  | go test -race     | Unit tests with race detection |
| Build | make build        | Compile binary to `bin/api`    |
