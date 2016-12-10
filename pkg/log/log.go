package log

import (
	"io"
	"log"
	"os"
)

type Level int

var logger *log.Logger

func Init(logLevel Level, out io.Writer) {
	logger = log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Fatal(msg string) {
	if logger == nil {
		Init(0, os.Stderr)
	}
	logger.Fatalln("FATAL - " + msg)
}

func Error(msg string) {
	if logger == nil {
		Init(0, os.Stderr)
	}
	logger.Println("ERROR - " + msg)
}

func Warning(msg string) {
	if logger == nil {
		Init(0, os.Stderr)
	}
	logger.Println("WARNING - " + msg)
}

func Info(msg string) {
	if logger == nil {
		Init(0, os.Stderr)
	}
	logger.Println("INFO - " + msg)
}

func Debug(msg string) {
	if logger == nil {
		Init(0, os.Stderr)
	}
	logger.Println("DEBUG - " + msg)
}

func Debugf(format string, v ...interface{}) {
	if logger == nil {
		Init(0, os.Stderr)
	}
	logger.Printf("DEBUG - "+format, v...)
}
