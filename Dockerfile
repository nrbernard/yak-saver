# Stage 1: Build frontend
FROM node:20-alpine AS frontend-builder
WORKDIR /frontend-build

# Copy package files and install dependencies
COPY frontend/package*.json ./
RUN npm ci

# Copy frontend source and build
COPY frontend/ ./
# Set empty API base URL for production (same origin)
ENV VITE_API_BASE_URL=
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.25-alpine AS backend-builder
WORKDIR /build

# Install build dependencies for CGO (required for SQLite)
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY sql/ ./sql/
COPY sqlc.yaml ./

# Build binary with CGO enabled for SQLite support
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd

# Stage 3: Runtime
FROM alpine:latest
WORKDIR /app

# Install runtime dependencies and Go (needed for goose)
RUN apk add --no-cache sqlite-libs ca-certificates go

# Install goose for database migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest && \
    mv /root/go/bin/goose /usr/local/bin/goose

# Copy backend binary
COPY --from=backend-builder /build/main ./

# Copy frontend static files
COPY --from=frontend-builder /frontend-build/dist ./frontend/dist

# Copy migration files
COPY sql/schema ./sql/schema

# Create data directory for volume mount
RUN mkdir -p /data

# Copy startup script
COPY docker-entrypoint.sh ./
RUN chmod +x docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./docker-entrypoint.sh"]
