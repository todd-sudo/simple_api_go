package app

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/todd-sudo/todo/internal/config"
	"github.com/todd-sudo/todo/internal/handler"
	"github.com/todd-sudo/todo/internal/repository"
	"github.com/todd-sudo/todo/pkg/server"

	"os"

	"github.com/todd-sudo/todo/internal/service"
	log "github.com/todd-sudo/todo/pkg/logger"
)

func Run() {

	cfg := config.GetConfig()
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Error(err)
	}

	log.Info("Connect to database successfully!")

	ctx := context.Background()

	repos := repository.NewRepository(ctx, db)
	services := service.NewService(ctx, *repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg.App.Port, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started on http://127.0.0.1:" + cfg.App.Port)

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
