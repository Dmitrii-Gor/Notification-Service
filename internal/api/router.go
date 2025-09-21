package api

import (
	"github.com/Dmitrii-Gor/notification-bot/internal/api/handlers"
	"github.com/Dmitrii-Gor/notification-bot/internal/config"
	"github.com/gin-gonic/gin"
)

func GinRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	emailGroup := r.Group("email")

	auth := handlers.NewAuthHandler()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	emailGroup.POST("/register")

	return r
}
