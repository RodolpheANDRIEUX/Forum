package initializer

import (
	"forum/Log"
	"github.com/natefinch/lumberjack"
	"log"
	"os"
)

var LogFile *lumberjack.Logger

func InitLogs() {

	logsDir := "./logs"
	// Create logs directory if it does not exist
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		err = os.MkdirAll(logsDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	LogFile = &lumberjack.Logger{
		Filename:   logsDir + "/gin.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28, // in days
		Compress:   false,
	}

	if err := LogFile.Rotate(); err != nil {
		Log.Err.Fatal(err)
	}

	Log.InitErrorLog(LogFile)
	log.SetOutput(LogFile)
}
