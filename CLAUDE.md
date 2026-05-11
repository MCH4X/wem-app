# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

WEM — Where's My Money? is a personal finance tracking application with a Go backend and React frontend. The app helps users track their expenses with the tagline "Track it before it's gone again."

## Architecture

The application follows a clean architecture pattern with clear separation of concerns:

### Backend (Go)
- **Entry Point**: `backend/cmd/main.go` - Initializes database and sets up routes
- **Database Layer**: 
  - `backend/internal/database/` - Database connection and migrations
  - Uses PostgreSQL with GORM as ORM
- **Domain Layer**:
  - `backend/internal/models/` - Data models (e.g., User)
  - `backend/internal/repositories/` - Data access layer
  - `backend/internal/services/` - Business logic layer
- **Presentation Layer**:
  - `backend/internal/handlers/` - HTTP request handlers
  - `backend/internal/routes/` - Route definitions using Gin

### Frontend (React + TypeScript)
- Built with Vite for fast development and building
- Uses React Compiler for performance optimization
- ESLint for code quality
- TypeScript for type safety

## Development Commands

### Backend
```bash
cd backend

# Development workflow (runs all checks)
make dev

# Individual commands:
make lint      # Run golangci-lint
make test     # Run tests with race detection
make build    # Build binary to bin/api
make run      # Run the server
make clean    # Remove build artifacts
```

### Frontend
```bash
cd frontend

# Development:
npm run dev    # Start dev server
npm run build  # Build for production
npm run lint   # Run ESLint
npm run preview # Preview production build
```

## Database Setup

The backend uses PostgreSQL configured in `backend/docker-compose.yml`:
```yaml
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 1234
      POSTGRES_DB: mydb
    ports: ["5432:5432"]
```

Start the database with:
```bash
cd backend
docker-compose up -d
```

## API Endpoints

Current endpoints (based on routes/router.go):
- `GET /users` - Get all users
- `POST /users` - Create a new user

## CI/CD

The backend CI pipeline runs on:
- Push to main branch
- Pull requests to main branch

The pipeline uses `dorny/paths-filter` to only run when backend files are modified, but always returns a status to avoid Branch Protection deadlocks.

Steps:
1. Lint with golangci-lint
2. Test with `go test -race -v`
3. Build with `make build`

## Project Structure Notes

- Frontend and backend are in separate directories
- No shared configuration or workspace setup
- Frontend uses pnpm (evidenced by pnpm-lock.yaml)
- Backend uses Go 1.26.2
- Frontend uses React 19.2.6 with TypeScript