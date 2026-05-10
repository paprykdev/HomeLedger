# =========================================
# FRONTEND BUILD
# =========================================

FROM node:22-alpine AS frontend-builder

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend .

RUN npm run build


# =========================================
# BACKEND BUILD
# =========================================

FROM golang:1.25-alpine AS backend-builder

WORKDIR /backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend .

RUN go build -o homeledger ./cmd/api


# =========================================
# FINAL IMAGE
# =========================================

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache \
    ca-certificates \
    nodejs \
    caddy \
    supervisor

# backend binary
COPY --from=backend-builder /backend/homeledger ./homeledger

# next standalone build
COPY --from=frontend-builder /frontend/.next/standalone ./frontend
COPY --from=frontend-builder /frontend/.next/static ./frontend/.next/static
COPY --from=frontend-builder /frontend/public ./frontend/public
COPY deploy/production/Caddyfile /app/Caddyfile
COPY deploy/production/supervisord.conf /etc/supervisord.conf


# storage
RUN mkdir -p /app/storage

ENV PORT=8080
ENV HOMELEDGER_DB_PATH=/app/storage/homeledger.db

EXPOSE 8080

CMD ["supervisord", "-c", "/etc/supervisord.conf"]
