# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

WEM — Where's My Money? is a personal finance tracking application with a Go backend and React frontend. The app helps users track their expenses with the tagline "Track it before it's gone again."

## Architecture

The application follows a clean architecture pattern with clear separation of concerns:

### Backend (Go)
- **Entry Point**: `backend/cmd/main.go` - Initializes database and sets up routes
- **Database Layer**: `backend/internal/database/db.go` - PostgreSQL connection built from env vars (`DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT`, `DB_SSLMODE`)
  - Uses PostgreSQL with GORM as ORM
  - No auto-migration in main.go; schema is managed separately
- **Domain Layer**:
  - `backend/internal/models/` - Data models (`user.model.go`: ID, Name, Age)
  - `backend/internal/repositories/` - Data access layer (`user.repository.go`)
  - `backend/internal/services/` - Business logic layer (`user.service.go`)
- **Presentation Layer**:
  - `backend/internal/handlers/` - HTTP request handlers (`user.handler.go`)
  - `backend/internal/routes/router.go` - Gin router on port 8080

### Frontend (React + TypeScript)
- **Entry**: `frontend/src/main.tsx` → `frontend/src/App.tsx`
- Built with Vite 8 for fast development and building
- Uses Babel React Compiler plugin (`@babel/plugin-react-compiler`) for performance optimization
- ESLint with TypeScript, React Hooks, and React Refresh plugins
- TypeScript 6 with ES2023 target, bundler module resolution
- Currently a starter template — no backend API integration yet

## Key Dependencies

### Backend
- Go 1.26.2, module: `github.com/MCH4X/wem-app/backend`
- `github.com/gin-gonic/gin` v1.12.0 — web framework
- `gorm.io/gorm` v1.31.1 + `gorm.io/driver/postgres` v1.6.0 — ORM

### Frontend
- React 19.2.6, TypeScript 6.0.2
- Vite 8.0.12
- Package manager: pnpm (pnpm-lock.yaml present; use `pnpm` not `npm`)

## Development Commands

### Backend
```bash
cd backend

make dev       # Lint + test + build (full workflow)
make lint      # Run golangci-lint
make test      # Run tests (no -race flag in Makefile — check before adding)
make build     # Build binary to bin/api
make run       # Run the server
make clean     # Remove build artifacts
```

### Frontend
```bash
cd frontend

pnpm dev       # Start dev server
pnpm build     # Build for production
pnpm lint      # Run ESLint
pnpm preview   # Preview production build
```

## Database Setup

PostgreSQL is configured in `backend/docker-compose.yml` using variable substitution from `backend/.env`:

```bash
cd backend
cp .env.example .env   # fill in real values
docker-compose up -d
```

Schema and seed:
```bash
make migration-up  # create tables (uses golang-migrate, files in backend/migrations/)
make seed          # insert sample data (skips if table is non-empty)
make migration-down  # roll back all migrations
```

## API Endpoints

Server listens on `:8080`. Current endpoints:
- `GET /users` — Get all users
- `POST /users` — Create a new user

## Linting

Backend linter config at `backend/.golangci.yml`:
- Enabled: govet, errcheck, staticcheck, unused, ineffassign
- Timeout: 5m

## CI/CD

Backend CI pipeline (`.github/workflows/backend.yml`):
- Triggers: push/PR to `main`, filtered to `backend/**` changes
- Uses path filtering to skip when only frontend files change, but always returns a status to avoid Branch Protection deadlocks
- Steps: Lint → Test (`go test -race`) → Build (`make build`)
- Runs on Node 24 GitHub Actions runner

## Git Commit Convention

This project follows **Conventional Commits** (`type: short description`).

| Type | When to use |
|------|-------------|
| `feat` | New feature or endpoint |
| `fix` | Bug fix |
| `refactor` | Code change that is neither a fix nor a feature |
| `docs` | Documentation only |
| `chore` | Config, deps, tooling, CI tweaks |
| `ci` | CI/CD pipeline changes |
| `test` | Adding or updating tests |
| `init` | Initial project/module setup |

**Rules:**
- Lowercase subject, no trailing period
- Imperative mood — "add user endpoint", not "added" or "adds"
- Keep the subject under 72 characters
- No ticket references in the subject line

```
feat: add expense category filter
fix: return 404 when user not found
chore: upgrade gin to v1.12
```

## Project Structure

```
wem-app/
├── .github/workflows/backend.yml
├── backend/
│   ├── cmd/
│   │   ├── main.go          # API server entry point
│   │   ├── migrate/main.go  # migration runner (up/down)
│   │   └── seed/main.go     # seed runner
│   ├── migrations/
│   │   ├── 000001_create_users_table.up.sql
│   │   └── 000001_create_users_table.down.sql
│   ├── internal/
│   │   ├── database/db.go
│   │   ├── handlers/user.handler.go
│   │   ├── models/user.model.go
│   │   ├── repositories/user.repository.go
│   │   ├── routes/router.go
│   │   └── services/user.service.go
│   ├── .env.example
│   ├── .golangci.yml
│   ├── docker-compose.yml
│   ├── go.mod / go.sum
│   └── Makefile
└── frontend/
    ├── public/
    ├── src/
    │   ├── App.tsx
    │   ├── App.css
    │   ├── main.tsx
    │   └── index.css
    ├── eslint.config.js
    ├── vite.config.ts
    ├── tsconfig.json / tsconfig.app.json / tsconfig.node.json
    └── package.json / pnpm-lock.yaml
```
