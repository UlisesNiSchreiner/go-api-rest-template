package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/config"
	"github.com/your-org/go-rest-layered-template/internal/httpserver"
	"github.com/your-org/go-rest-layered-template/internal/logger"
	"github.com/your-org/go-rest-layered-template/internal/platform/db"
	"github.com/your-org/go-rest-layered-template/internal/repositories/mysqlrepo"
	"github.com/your-org/go-rest-layered-template/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg.Env)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sqlDB, err := db.NewMySQL(ctx, cfg.MySQL, log)
	if err != nil {
		log.Fatal("database connection failed", logger.Err(err))
	}
	defer func() { _ = sqlDB.Close() }()

	userRepo := mysqlrepo.NewUserRepository(sqlDB)
	userSvc := services.NewUserService(userRepo)

	app := httpserver.New(cfg, log, userSvc)

	server := &http.Server{
		Addr:              cfg.HTTP.Addr(),
		Handler:           app.Router(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Info("http server started", logger.String("addr", cfg.HTTP.Addr()))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("http server failed", logger.Err(err))
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("shutting down")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("shutdown error", logger.Err(err))
	}
	log.Info("bye")
}
