-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks ADD COLUMN completed_at DATETIME;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks DROP COLUMN completed_at;
-- +goose StatementEnd
