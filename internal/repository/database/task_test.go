package storage_test

/*
import (
	"backend-todo-list/services/task/internal/models"
	"context"
	"testing"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockDB struct {
	mock.Mock
}

func (m *MockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	ret := m.Called(ctx, sql, args)
	return ret.Get(0).(pgx.Row)
}

func (m *MockDB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	ret := m.Called(ctx, sql, args)
	return ret.Get(0).(pgconn.CommandTag), ret.Error(1)
}

func (m *MockDB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	ret := m.Called(ctx, sql, args)
	return ret.Get(0).(pgx.Rows), ret.Error(1)
}


func TestCreateTask(t *testing.T) {
	mockDB := new(MockDB)
	storage := storage.Storage{db: mockDB}

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now(),
		Completed:   false,
	}

	mockDB.On("QueryRow", mock.Anything, "INSERT INTO tasks (title, description, due_date, completed) VALUES ($1, $2, $3, $4) RETURNING id", task.Title, task.Description, task.DueDate, task.Completed).
		Return(1, nil)

	id, err := storage.CreateTask(context.Background(), task)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockDB.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockDB := new(MockDB)
	storage := storage.Storage{db: mockDB}

	taskID := 1
	expectedTask := models.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now(),
		Completed:   false,
	}

	mockDB.On("QueryRow", mock.Anything, "SELECT id, title, description, due_date, completed FROM tasks WHERE id = $1", taskID).
		Return(expectedTask, nil)

	task, err := storage.GetTask(context.Background(), taskID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, *task)
	mockDB.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockDB := new(MockDB)
	storage := storage.Storage{db: mockDB}

	task := models.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated description",
		DueDate:     time.Now(),
		Completed:   true,
	}

	mockDB.On("Exec", mock.Anything, "UPDATE tasks SET title = $1, description = $2, due_date = $3, completed = $4 WHERE id = $5",
		task.Title, task.Description, task.DueDate, task.Completed, task.ID).Return(nil)

	err := storage.UpdateTask(context.Background(), task)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockDB := new(MockDB)
	storage := storage.Storage{db: mockDB}

	taskID := 1

	mockDB.On("Exec", mock.Anything, "DELETE FROM tasks WHERE id = $1", taskID).Return(nil)

	err := storage.DeleteTask(context.Background(), taskID)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestListTasks(t *testing.T) {
	mockDB := new(MockDB)
	storage := storage.Storage{db: mockDB}

	expectedTasks := []models.Task{
		{ID: 1, Title: "Task 1", Description: "First Task", DueDate: time.Now(), Completed: false},
		{ID: 2, Title: "Task 2", Description: "Second Task", DueDate: time.Now(), Completed: true},
	}

	rows := pgxmock.NewRows([]string{"id", "title", "description", "due_date", "completed"}).
		AddRow(expectedTasks[0].ID, expectedTasks[0].Title, expectedTasks[0].Description, expectedTasks[0].DueDate, expectedTasks[0].Completed).AddRow(expectedTasks[1].ID, expectedTasks[1].Title, expectedTasks[1].Description, expectedTasks[1].DueDate, expectedTasks[1].Completed)

	mockDB.On("Query", mock.Anything, "SELECT id, title, description, due_date, completed FROM tasks WHERE completed = $1 ORDER BY due_date LIMIT $2 OFFSET $3", false, 10, 0).
		Return(rows, nil)

	tasks, err := storage.ListTasks(context.Background(), false, 10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, tasks)
	mockDB.AssertExpectations(t)
}

func TestGetTasksByDate(t *testing.T) {
	mockDB := new(MockDB)
	storage := storage.Storage{db: mockDB}

	date := time.Now()
	expectedTasks := []models.Task{
		{ID: 1, Title: "Task 1", Description: "First Task", DueDate: date, Completed: false},
	}

	rows := pgxmock.NewRows([]string{"id", "title", "description", "due_date", "completed"}).
		AddRow(expectedTasks[0].ID, expectedTasks[0].Title, expectedTasks[0].Description, expectedTasks[0].DueDate, expectedTasks[0].Completed)

	mockDB.On("Query", mock.Anything,
		"SELECT id, title, description, due_date, completed FROM tasks WHERE due_date::date = $1 AND completed = $2",
		date.Format("2006-01-02"), false).Return(rows, nil)

	tasks, err := storage.GetTasksByDate(context.Background(), date, false)

	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, tasks)
	mockDB.AssertExpectations(t)
}
*/