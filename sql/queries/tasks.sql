-- name: GetTasksOrdered :many
SELECT id, project_id, parent_task_id, content, link, created_at, updated_at 
FROM tasks 
ORDER BY project_id, CASE WHEN parent_task_id IS NULL THEN 0 ELSE 1 END, id;
