package src

import (
	"file-server/src/handler"
)

func configRouter() {
	App.GET("/download/:filename", handler.DownloadFile)
	App.GET("/download/all", handler.DownloadFile)
	App.POST("/upload", handler.UploadFileHandler)
}
