package storage

import (
	"backend-todo-list/services/task/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) CreateTask(ctx context.Context, task models.Task) (int, error) {
	var id int
	err := s.db.QueryRow(ctx,
		"INSERT INTO tasks (title, description, due_date, completed) VALUES ($1, $2, $3, $4) RETURNING id",
		task.Title, task.Description, task.DueDate, task.Completed).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}
	return id, nil
}

func (s *Storage) GetTask(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task
	err := s.db.QueryRow(ctx,
		"SELECT id, title, description, due_date, completed FROM tasks WHERE id = $1", id).Scan(
		&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("task not found: %d", id)
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return &task, nil
}

func (s *Storage) UpdateTask(ctx context.Context, task models.Task) error {
	_, err := s.db.Exec(ctx,
		"UPDATE tasks SET title = $1, description = $2, due_date = $3, completed = $4 WHERE id = $5",
		task.Title, task.Description, task.DueDate, task.Completed, task.ID)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (s *Storage) DeleteTask(ctx context.Context, id int) error {
	_, err := s.db.Exec(ctx,
		"DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (s *Storage) ListTasks(ctx context.Context, completed bool, limit int, offset int) ([]models.Task, error) {
	rows, err := s.db.Query(ctx,
		"SELECT id, title, description, due_date, completed FROM tasks WHERE completed = $1 ORDER BY due_date LIMIT $2 OFFSET $3",
		completed, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Storage) GetTasksByDate(ctx context.Context, date time.Time, completed bool) ([]models.Task, error) {
	rows, err := s.db.Query(ctx,
		"SELECT id, title, description, due_date, completed FROM tasks WHERE due_date::date = $1 AND completed = $2",
		date.Format("2006-01-02"), completed)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by date: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
