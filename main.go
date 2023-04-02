package main

import (
	"file-server/src"
	"file-server/src/config"
	"file-server/src/db"
	"os"

	"go.uber.org/zap"
)

func main() {
	_, err := db.GcsClient()
	if err != nil {
		config.LOGGER.Panic("Error connecting to GCS. ", zap.Error(err))
	}
	_, errMdb := db.MongoClient()
	if err != nil {
		config.LOGGER.Panic("Error connectiing to MongoDB. ", zap.Error(errMdb))
	}

	src.InitApp()
	src.App.Run(os.Getenv("PORT"))
}
