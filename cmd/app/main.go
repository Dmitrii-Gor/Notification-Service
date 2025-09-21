package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/Dmitrii-Gor/notification-bot/internal/api"
	"github.com/Dmitrii-Gor/notification-bot/internal/config"
	"github.com/Dmitrii-Gor/notification-bot/internal/storage"
	"github.com/Dmitrii-Gor/notification-bot/pkg/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const httpAddr = ":8080"

func main() {
	cfg := config.Load()

	logger.InitLogger(cfg.Environment)
	defer logger.Sync()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Error("open database", zap.Error(err))
		panic(err)
	}

	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(cfg.DBTimeout)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DBTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logger.Error("ping database", zap.Error(err))
		panic(err)
	}

	defer db.Close()

	userRepo := storage.NewUserRepo(db)

	router := api.GinRouter(cfg, userRepo)

	logger.Info("server starting", zap.String("addr", httpAddr), zap.String("env", cfg.Environment))
	if err := router.Run(httpAddr); err != nil {
		logger.Error("server shutdown", zap.Error(err))
	}
}
