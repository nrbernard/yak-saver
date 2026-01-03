-- name: GetTasksOrdered :many
SELECT id, project_id, parent_task_id, content, link, created_at, updated_at 
FROM tasks 
ORDER BY project_id, CASE WHEN parent_task_id IS NULL THEN 0 ELSE 1 END, id;

-- name: UpdateTask :exec
UPDATE tasks 
SET content = ?1, link = ?2, updated_at = CURRENT_TIMESTAMP 
WHERE id = ?3;

-- name: CreateTask :one
INSERT INTO tasks (project_id, parent_task_id, content, link)
VALUES (?1, ?2, ?3, ?4)
RETURNING id, project_id, parent_task_id, content, link, created_at, updated_at;

-- name: GetTasksByParentID :many
SELECT id, project_id, parent_task_id, content, link, created_at, updated_at
FROM tasks
WHERE parent_task_id = ?1;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?1;