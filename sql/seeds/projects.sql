-- Idempotent seed data for projects table
INSERT INTO projects (name)
SELECT 'Dining Room'
WHERE NOT EXISTS (SELECT 1 FROM projects WHERE name = 'Dining Room');

INSERT INTO projects (name)
SELECT 'Kitchen Renovation'
WHERE NOT EXISTS (SELECT 1 FROM projects WHERE name = 'Kitchen Renovation');

INSERT INTO projects (name)
SELECT 'Home Office'
WHERE NOT EXISTS (SELECT 1 FROM projects WHERE name = 'Home Office');

