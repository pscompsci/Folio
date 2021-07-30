package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Err
	Fatal
)

func (l Level) String() string {
	return [...]string{"DEBUG", "INFO ", "WARN ", "ERROR", "FATAL"}[l]
}

type Logger interface {
	SetLevel(level string)
	SetTimeFormat(timeformat string)
	LogDebug(message string)
	LogInformation(message string)
	LogWarning(message string)
	LogError(mesage string)
	LogFatal(message string)
}

type FileLogger struct {
	filename   string
	mutex      *sync.Mutex
	level      Level
	timeformat string
}

// NewFileLogger returns a Logger that logs to file
func NewFileLogger(filename string, timeformat string) *FileLogger {
	fl := &FileLogger{
		filename:   filename,
		mutex:      &sync.Mutex{},
		level:      Info,
		timeformat: timeformat,
	}
	return fl
}

// SetLevel sets the minimum logging level (defaul = INFORMATION)
func (fl *FileLogger) SetLevel(level string) {
	switch level {
	case "DEBUG":
		fl.level = Debug
	case "INFORMATION":
		fl.level = Info
	case "WARNING":
		fl.level = Warn
	case "ERROR":
		fl.level = Err
	case "FATAL":
		fl.level = Fatal
	default:
		fl.level = Info
	}
}

// SetTimeFormat sets the timeformat for the log messages
func (fl *FileLogger) SetTimeFormat(timeformat string) {
	fl.timeformat = timeformat
}

// LogDebug writes a debug message to the log file
func (fl *FileLogger) LogDebug(msg string) {
	fl.logToFile(Debug, msg)
}

// LogInformation writes an information message to the log file
func (fl *FileLogger) LogInformation(msg string) {
	if fl.checkLevel(Info) {
		fl.logToFile(Info, msg)
	}
}

// LogWarning logs a warning message to the log file
func (fl *FileLogger) LogWarning(msg string) {
	if fl.checkLevel(Warn) {
		fl.logToFile(Warn, msg)
	}
}

// LogError logs an error message to the log file
func (fl *FileLogger) LogError(msg string) {
	if fl.checkLevel(Err) {
		fl.logToFile(Err, msg)
	}
}

// LogFatal logs a fatal message to the log file
func (fl *FileLogger) LogFatal(msg string) {
	if fl.checkLevel(Fatal) {
		fl.logToFile(Fatal, msg)
	}
}

func (fl *FileLogger) checkLevel(level Level) bool {
	if level < Debug || level > Fatal {
		return false
	}
	return fl.level <= level
}

func (fl *FileLogger) logToFile(level Level, message string) {
	fl.mutex.Lock()
	defer fl.mutex.Unlock()

	msg := formatMessage(time.Now(), fl.timeformat, level.String(), message)

	f, err := os.OpenFile(fl.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("logger: Could not open or create log file\n%s\n", msg))
	}

	defer func() {
		err := f.Close()
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("logger: Closing failed on log file: %s\n", fl.filename))
			return
		}
	}()

	log(f, msg)
}

func formatMessage(time time.Time, timeformat string, level string, message string) string {
	return fmt.Sprintf("%s %s: %s\n", time.Format(timeformat), level, message)
}

func log(logFile io.Writer, fmtMsg string) error {
	if _, err := logFile.Write([]byte(fmtMsg)); err != nil {
		return err
	}
	return nil
}
