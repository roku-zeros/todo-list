package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend-todo-list/services/task/internal/app"
)

const (
	serviceName = "market"
)

func main() {
	path := os.Getenv("CONFIG")
	if path == "" {
		path = "config.yaml"
	}

	createCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	app, err := app.New(createCtx, path)
	if err != nil {
		panic(err)
	}

	runCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()
	go func() {
		err := app.Run(context.Background())
		if err != nil {
			log.Fatalf("Error during app run: %+v", err)
		}
	}()
	<-runCtx.Done()

	stopCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err = app.Stop(stopCtx); err != nil {
		panic(err)
	}
}
