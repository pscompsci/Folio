package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	debug = iota
	info
	warn
	err
	fatal
)

type Logger struct {
	filename   string
	mutex      *sync.Mutex
	level      int
	timeformat string
}

// NewLogger returns a new Logger instance
func NewLogger(filename string, timeformat string) *Logger {
	l := &Logger{
		filename:   filename,
		mutex:      &sync.Mutex{},
		level:      info,
		timeformat: timeformat,
	}
	return l
}

// SetTimeFormat sets the timeformat for the log messages
func (l *Logger) SetTimeFormat(timeformat string) {
	l.timeformat = timeformat
}

// SetLevel sets the minimum logging level (defaul = INFORMATION)
func (l *Logger) SetLevel(level string) {
	switch level {
	case "DEBUG":
		l.level = debug
	case "INFORMATION":
		l.level = info
	case "WARNING":
		l.level = warn
	case "ERROR":
		l.level = err
	case "FATAL":
		l.level = fatal
	default:
		l.level = info
	}
}

// LogDebug writes a debug message to the log file
func (l *Logger) LogDebug(msg string) {
	l.log(debug, msg)
}

// LogInformation writes an information message to the log file
func (l *Logger) LogInformation(msg string) {
	l.log(info, msg)
}

// LogWarning logs a warning message to the log file
func (l *Logger) LogWarning(msg string) {
	l.log(warn, msg)
}

// LogError logs an error message to the log file
func (l *Logger) LogError(msg string) {
	l.log(err, msg)
}

// LogFatal logs a fatal message to the log file
func (l *Logger) LogFatal(msg string) {
	l.log(fatal, msg)
}

func (l *Logger) log(level int, message string) {
	if level < l.level {
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	fmtMsg := fmt.Sprintf("%s %s: %s\n", time.Now().Format(l.timeformat), levelToString(level), message)

	logFile, err := os.OpenFile(l.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("logger: %s\n", fmtMsg))
	}

	defer func() {
		err := logFile.Close()
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("logger: Closing failed on log file: %s\n", l.filename))
			return
		}
	}()

	if _, err := logFile.WriteString(fmtMsg); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("logger: Writing failed on log file: %s\n", l.filename))
		return
	}

}

func levelToString(level int) string {
	switch level {
	case debug:
		return "DEBUG"
	case info:
		return "INFO "
	case warn:
		return "WARN "
	case err:
		return "ERROR"
	case fatal:
		return "FATAL"
	default:
		return "INFO "
	}
}
