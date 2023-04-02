package config

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	HOST       = os.Getenv("HOST")
	PORT       = os.Getenv("PORT")
	URL        = fmt.Sprintf("http://%s:%s", HOST, PORT)
	LOGGER     *zap.Logger
	loggerExec sync.Once
)

func Logger() *zap.Logger {
	loggerExec.Do(func() {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		defaultLogLevel := zapcore.DebugLevel
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)

		LOGGER = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})

	return LOGGER
}
