.PHONY: migrate seed generate dev-backend dev-frontend dev

migrate:
	goose -dir sql/schema sqlite3 data/yak-saver.db up

seed:
	@echo "Seeding database..."
	@sqlite3 data/yak-saver.db < sql/seeds/projects.sql

generate:
	sqlc generate

dev-backend:
	@echo "Starting backend with hot reload..."
	air

dev-frontend:
	@echo "Starting frontend dev server..."
	cd frontend && npm run dev

dev:
	@echo "Starting Yak Saver (backend + frontend)..."
	@echo "Backend: http://localhost:8080"
	@echo "Frontend: http://localhost:5173"
	@echo ""
	@trap 'kill 0' INT; \
	air & \
	cd frontend && npm run dev & \
	wait