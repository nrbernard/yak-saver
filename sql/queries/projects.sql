-- name: GetProjects :many
SELECT id, name, created_at, updated_at FROM projects ORDER BY created_at DESC;

-- name: CreateProject :one
INSERT INTO projects (name)
VALUES (?1)
RETURNING id, name, created_at, updated_at;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = ?1;
