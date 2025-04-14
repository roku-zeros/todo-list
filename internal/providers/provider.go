package providers

import (
	"backend-todo-list/services/task/internal/models"
	"context"
	"time"
)

type Storage interface {
	CreateTask(ctx context.Context, task models.Task) (int, error)
	GetTask(ctx context.Context, id int) (*models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) error
	DeleteTask(ctx context.Context, id int) error
	ListTasks(ctx context.Context, completed bool, limit int, offset int) ([]models.Task, error)
	GetTasksByDate(ctx context.Context, date time.Time, completed bool) ([]models.Task, error)
}

type TaskProvider struct {
	storage Storage
}

func NewTaskService(storage Storage) TaskProvider {
	return TaskProvider{
		storage: storage,
	}
}
