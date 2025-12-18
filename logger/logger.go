package logger

import (
	"log"
	"os"
	"strings"
)

// Logger is a simple wrapper around Go's log package.
// It gives us readable log levels like INFO, DEBUG, WARN, ERROR.
type Logger struct {
	level string

	infoLogger  *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger creates and returns a new logger instance.
func NewLogger(level string) *Logger {
	level = strings.ToLower(level)

	return &Logger{
		level: level,

		infoLogger:  log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "[WARN]  ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
	}
}

// Debug logs extra details (shown only in debug mode)
func (l *Logger) Debug(message string) {
	if l.level == "debug" {
		l.debugLogger.Println(message)
	}
}

// Info logs normal application flow
func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

// Warn logs something unexpected but not fatal
func (l *Logger) Warn(message string) {
	l.warnLogger.Println(message)
}

// Error logs serious problems
func (l *Logger) Error(message string) {
	l.errorLogger.Println(message)
}
