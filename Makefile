.PHONY: migrate

migrate:
	goose -dir sql/schema sqlite3 data/yak-saver.db up

generate:
	sqlc generate