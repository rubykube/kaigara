package log

import (
	"io"
	"log"
)

type Level int

var logger *log.Logger

func Init(logLevel Level, out io.Writer) {
	logger = log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Fatal(msg string) {
	logger.Fatalln("FATAL - " + msg)
}

func Error(msg string) {
	logger.Println("ERROR - " + msg)
}

func Warning(msg string) {
	logger.Println("WARNING - " + msg)
}

func Info(msg string) {
	logger.Println("INFO - " + msg)
}

func Debug(msg string) {
	logger.Println("DEBUG - " + msg)
}

func Debugf(format string, v ...interface{}) {
	logger.Printf("DEBUG - "+format, v...)
}
