# HomeLedger

Privacy-first self-hosted finance tracker.

## Features

- Self-hosted
- SQLite
- REST API
- Frontend + backend Docker setup
- Transaction tracking

## Tech Stack

- Backend: Go + Chi + SQLite
- Frontend: Next.js + React + pnpm
- Docker / Docker Compose

## Local development (without Docker)

### Backend

```bash
cd backend
go run ./cmd/migrate
PORT=8080 go run ./cmd/api
```

Backend API: `http://localhost:8080/api`

### Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

Frontend runs on `http://localhost:3000` by default.

## Docker Compose (development)

Use the root compose file for development (bind mounts + dev commands):

```bash
docker compose up
```

Services:
- Frontend: `http://localhost:3000`
- Backend API: `http://localhost:8080/api`

## Docker Compose (production image)

```bash
docker compose -f docker-compose.prod.yml up -d
```

This uses the published image `paprykdev/homeledger:latest`.

## Environment variables

Backend:
- `PORT` (default: `3000`)
- `JWT_SECRET` (required for real deployments)
- `HOMELEDGER_DB_PATH` (default path is handled in code)

Frontend:
- `NEXT_PUBLIC_API_BASE_URL` (used in container setup)

## API reference

See `backend/API_ENDPOINTS.md` for frontend-oriented endpoint documentation.
