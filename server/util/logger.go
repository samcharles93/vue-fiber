package util

import (
	"os"

	"github.com/rs/zerolog"
)

var file *os.File
var logger zerolog.Logger

func InitLogger() (zerolog.Logger, error) {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return zerolog.Logger{}, err
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(file, consoleWriter)
	logger = zerolog.New(multi).With().Timestamp().Logger()
	logger = logger.Level(zerolog.DebugLevel)

	return logger, nil
}

func CloseLogger() {
	if file != nil {
		file.Close()
	}
}
