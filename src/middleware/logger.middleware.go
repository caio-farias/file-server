package middleware

import (
	"file-server/src/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddelware(c *gin.Context) {
	config.LOGGER.Info(">> ",
		zap.String("method: ", c.Request.Method),
		zap.String("path: ", c.Request.URL.Path),
		zap.Any("body:", c.Request.Body),
	)

	c.Next()
}
