-- name: GetProjects :many
SELECT id, name, created_at, updated_at FROM projects;
