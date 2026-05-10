FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api && \
    CGO_ENABLED=0 GOOS=linux go build -o /out/migrate ./cmd/migrate

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/api /app/api
COPY --from=builder /out/migrate /app/migrate
COPY --from=builder /src/internal/database/migrations /app/internal/database/migrations

RUN mkdir -p /app/storage

ENV PORT=8080
ENV HOMELEDGER_DB_PATH=/app/storage/homeledger.db

EXPOSE 8080

CMD ["/bin/sh", "-c", "/app/migrate && /app/api"]
