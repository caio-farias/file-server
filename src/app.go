package src

import (
	"file-server/src/middleware"

	"github.com/gin-gonic/gin"
)

var App *gin.Engine = gin.Default()

func InitApp() {
	App.Static("/files", "./storage")
	App.MaxMultipartMemory = 8 << 20 // 8 MiB
	App.Use(middleware.LoggerMiddelware)
	configRouter()
}
