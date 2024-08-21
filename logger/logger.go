package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"todoList/configs"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func Init() error {
	logParams := configs.AppSettings.LogParams

	if _, err := os.Stat(logParams.LogDirectory); os.IsNotExist(err) {
		err = os.Mkdir(logParams.LogDirectory, 0755)
		if err != nil {
			return err
		}
	}

	// Инициализация логгеров lumberjack
	lumberLogInfo := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogInfo),
		MaxSize:    logParams.MaxSizeMegabytes, // мегабайты
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   // дни
		Compress:   logParams.Compress, // отключено по умолчанию
		LocalTime:  logParams.LocalTime,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogError),
		MaxSize:    logParams.MaxSizeMegabytes, // мегабайты
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   // дни
		Compress:   logParams.Compress, // отключено по умолчанию
		LocalTime:  logParams.LocalTime,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogWarn),
		MaxSize:    logParams.MaxSizeMegabytes, // мегабайты
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   // дни
		Compress:   logParams.Compress, // отключено по умолчанию
		LocalTime:  logParams.LocalTime,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogDebug),
		MaxSize:    logParams.MaxSizeMegabytes, // мегабайты
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   // дни
		Compress:   logParams.Compress, // отключено по умолчанию
		LocalTime:  logParams.LocalTime,
	}

	// Инициализация глобальных логгеров
	Info = log.New(gin.DefaultWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lumberLogError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(lumberLogWarn, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lumberLogDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	return nil
}
