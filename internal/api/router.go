package api

import (
	"github.com/Dmitrii-Gor/notification-bot/internal/api/handlers"
	"github.com/Dmitrii-Gor/notification-bot/internal/config"
	"github.com/Dmitrii-Gor/notification-bot/internal/domain"
	"github.com/gin-gonic/gin"
)

func GinRouter(cfg *config.Config, users domain.UserRepo) *gin.Engine {
	r := gin.Default()

	auth := handlers.NewAuthHandler(
		users,
		[]byte(cfg.JWT.Secret),
		cfg.JWT.AccessTTL,
		cfg.JWT.RefreshTTL,
		cfg.DBTimeout,
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	emailGroup := r.Group("email")
	emailGroup.POST("/register", auth.RegisterWithEmail)

	return r
}
