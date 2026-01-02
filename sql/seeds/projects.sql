-- Idempotent seed data for projects table
INSERT INTO projects (id, name)
SELECT 1, 'Dining Room'
WHERE NOT EXISTS (SELECT 1 FROM projects WHERE id = 1);

-- Idempotent seed data for tasks table
-- Top-level task: Art supply storage
INSERT INTO tasks (id, project_id, parent_task_id, content, link)
SELECT 2, 1, NULL, 'Art supply storage', NULL
WHERE NOT EXISTS (SELECT 1 FROM tasks WHERE id = 2);

-- Child task: Ikea Trofast
INSERT INTO tasks (id, project_id, parent_task_id, content, link)
SELECT 3, 1, 2, 'Ikea Trofast', 'https://www.ikea.com/us/en/p/trofast-storage-combination-with-boxes-light-white-stained-pine-light-orange-s09574857/#content'
WHERE NOT EXISTS (SELECT 1 FROM tasks WHERE id = 3);
