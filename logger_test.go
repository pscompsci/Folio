package main

import (
	"testing"
	"time"
)

func TestNewLoggerCreatesOK(t *testing.T) {
	l := NewLogger("log_test.log", time.RFC3339)
	level := l.level
	if level != info {
		t.Fatalf("Logger level incorrect. Want: 1, Got: %d", level)
	}
}

func TestSetLogLevelOK(t *testing.T) {
	l := NewLogger("log_test.log", time.RFC3339)
	l.SetLevel("FATAL")
	level := l.level
	if level != fatal {
		t.Fatalf("Logger level incorrect. Want: 4, Got: %d", level)
	}
}

func TestSetLogLevelDefaultsToInfoOnInvalidLevel(t *testing.T) {
	l := NewLogger("log_test.log", time.RFC3339)
	l.SetLevel("NOT_A_REAL_LEVEL")
	level := l.level
	if level != info {
		t.Fatalf("Logger level incorrect. Want: 1, Got: %d", level)
	}
}
