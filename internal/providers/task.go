package providers

import (
	"backend-todo-list/services/task/internal/models"
	"context"
	"time"
)

func (p *TaskProvider) CreateTask(ctx context.Context, task models.Task) (int, error) {
	return p.storage.CreateTask(ctx, models.Task{
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Completed:   task.Completed,
	})
}

func (p *TaskProvider) GetTask(ctx context.Context, id int) (*models.Task, error) {
	task, err := p.storage.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (p *TaskProvider) UpdateTask(ctx context.Context, task models.Task) error {
	return p.storage.UpdateTask(ctx, models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Completed:   task.Completed,
	})
}

func (p *TaskProvider) DeleteTask(ctx context.Context, id int) error {
	return p.storage.DeleteTask(ctx, id)
}

func (p *TaskProvider) ListTasks(ctx context.Context, completed bool, limit int, offset int) ([]models.Task, error) {
	tasks, err := p.storage.ListTasks(ctx, completed, limit, offset)
	if err != nil {
		return nil, err
	}

	var result []models.Task
	for _, t := range tasks {
		result = append(result, models.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			DueDate:     t.DueDate,
			Completed:   t.Completed,
		})
	}
	return result, nil
}

func (p *TaskProvider) GetTasksByDate(ctx context.Context, date time.Time, completed bool) ([]models.Task, error) {
	tasks, err := p.storage.GetTasksByDate(ctx, date, completed)
	if err != nil {
		return nil, err
	}

	var result []models.Task
	for _, t := range tasks {
		result = append(result, models.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			DueDate:     t.DueDate,
			Completed:   t.Completed,
		})
	}
	return result, nil
}
