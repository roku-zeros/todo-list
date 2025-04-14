-- Удаляем индексы, если они существуют
DROP INDEX IF EXISTS idx_tasks_due_date;
DROP INDEX IF EXISTS idx_tasks_is_completed;

-- Удаляем таблицу задач, если она существует
DROP TABLE IF EXISTS tasks;