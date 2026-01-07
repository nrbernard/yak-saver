#!/bin/sh
set -e

echo "Running database migrations..."
goose -dir /app/sql/schema sqlite3 ${DATABASE_PATH} up

echo "Starting application..."
exec ./main
