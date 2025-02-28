package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Debug(format string, args ...interface{})
}

type logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

func NewLogger() Logger {
	return &logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		warnLogger:  log.New(os.Stdout, "[WARN] ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
	}
}

func (l *logger) Info(format string, args ...interface{}) {
	message := fmt.Sprintf("%s %s", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprintf(format, args...))
	l.infoLogger.Println(message)
}

func (l *logger) Warn(format string, args ...interface{}) {
	message := fmt.Sprintf("%s %s", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprintf(format, args...))
	l.warnLogger.Println(message)
}

func (l *logger) Error(format string, args ...interface{}) {
	message := fmt.Sprintf("%s %s", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprintf(format, args...))
	l.errorLogger.Println(message)
}

func (l *logger) Debug(format string, args ...interface{}) {
	message := fmt.Sprintf("%s %s", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprintf(format, args...))
	l.debugLogger.Println(message)
}
