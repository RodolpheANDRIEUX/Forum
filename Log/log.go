package Log

import (
	"io"
	"log"
	"os"
)

var Err *log.Logger

func InitErrorLog(logFile io.Writer) {
	Err = log.New(io.MultiWriter(logFile, os.Stderr), "ERROR: ", log.LstdFlags)
}
