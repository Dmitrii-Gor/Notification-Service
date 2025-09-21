package main

import (
	"github.com/Dmitrii-Gor/notification-bot/internal/api"
	"github.com/Dmitrii-Gor/notification-bot/internal/config"
	"github.com/Dmitrii-Gor/notification-bot/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger("dev") // или "prod"
	defer logger.Sync()

	cfg := config.Load()

	gin := api.GinRouter(cfg)

	logger.Info("Server starting...", zap.String("addr", ":8080"))
	logger.Debug("debug log example")

	if err := gin.Run(":8080"); err != nil {
		panic(err)
	}
}
