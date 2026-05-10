# HomeLedger

Privacy-first self-hosted finance tracker.

## Features

- self-hosted
- SQLite
- REST API
- Docker support
- transaction tracking

## Tech Stack

- Go
- Chi
- SQLite
- Docker (planned)

## Getting Started

```bash
cd backend
go run ./cmd/api
```

Migrations are applied automatically on API startup.

Run migrations manually:

```bash
cd backend
go run ./cmd/migrate
```

## Docker Compose

Run frontend + backend together:

```bash
docker compose up --build
```

Services:
- Frontend: `http://localhost:3000`
- Backend API: `http://localhost:8080`

## Docker Compose (development)

Run both services in development mode with bind mounts:

```bash
docker compose -f docker-compose.dev.yml up
```

This runs:
- Backend via `go run` (with migrations on startup)
- Frontend via `pnpm dev` on `0.0.0.0:3000`
