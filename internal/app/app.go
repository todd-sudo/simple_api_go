package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/todd-sudo/todo/internal/handler"
	"github.com/todd-sudo/todo/internal/model"
	"github.com/todd-sudo/todo/internal/repository"
	"github.com/todd-sudo/todo/pkg/server"

	"github.com/todd-sudo/todo/internal/service"
	log "github.com/todd-sudo/todo/pkg/logger"
)

func Run() {
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Error(err)
	}

	db.AutoMigrate(&model.Item{}, &model.User{})

	log.Info("Connect to database successfully!")

	repos := repository.NewRepository(db)
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer("8000", handlers.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started on http://127.0.0.1:8000")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Info("Server stopped")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}
}
