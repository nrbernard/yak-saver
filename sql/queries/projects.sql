-- name: GetProjects :many
SELECT id, name, created_at, updated_at FROM projects;

-- name: CreateProject :one
INSERT INTO projects (name)
VALUES (?1)
RETURNING id, name, created_at, updated_at;
