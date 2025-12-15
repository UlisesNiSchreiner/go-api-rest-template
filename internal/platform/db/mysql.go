package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/config"
	"github.com/your-org/go-rest-layered-template/internal/logger"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(ctx context.Context, cfg config.MySQL, log *logger.Logger) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, err
	}

	log.Info("database connected")
	return db, nil
}
