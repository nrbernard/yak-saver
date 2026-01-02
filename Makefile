.PHONY: migrate seed generate

migrate:
	goose -dir sql/schema sqlite3 data/yak-saver.db up

seed:
	@echo "Seeding database..."
	@sqlite3 data/yak-saver.db < sql/seeds/projects.sql

generate:
	sqlc generate