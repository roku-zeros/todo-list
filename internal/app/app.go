package app

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"backend-todo-list/lib/config_utils"
	"backend-todo-list/lib/logger"
	"backend-todo-list/services/task/internal/config"
	"backend-todo-list/services/task/internal/providers"
	storage "backend-todo-list/services/task/internal/repository/database"
	"backend-todo-list/services/task/internal/server"
)

const (
	serviceName = "task"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type Storage interface {
	Stop(context.Context) error
}

type App struct {
	server  Server
	storage Storage
}

func New(ctx context.Context, configPath string) (*App, error) {

	config := config.Parse(configPath)

	err := logger.Init(config.Logger)
	if err != nil {
		return nil, err
	}

	postgresUser := config_utils.ReadSecret("POSTGRES_USER_FILE")
	postgresPassword := config_utils.ReadSecret("POSTGRES_PASSWORD_FILE")

	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s", postgresUser, postgresPassword, config.Postgres.Host, config.Postgres.Port)

	storage, err := storage.NewStorage(ctx, postgresUrl)
	if err != nil {
		logger.Error("db init error", zap.Error(err))
		return nil, err
	}

	taskProvider := providers.NewTaskService(storage)

	server := server.New(taskProvider)
	mux := http.NewServeMux()
	server.RegisterRoutes(mux)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &App{
		server:  httpServer,
		storage: storage,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	logger.Info("starting serve")
	return a.server.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) error {
	logger.Info("stopping at xserve")
	a.server.Shutdown(ctx)
	return a.storage.Stop(ctx)
}
