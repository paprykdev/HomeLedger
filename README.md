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
