package middleware

import (
	"github.com/Dmitrii-Gor/notification-bot/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic recovered", zap.Any("error", r))
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
