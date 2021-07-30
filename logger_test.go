package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestNewFileLogger(t *testing.T) {
	fl := NewFileLogger("log_test.log", time.RFC3339)
	level := fl.level
	if level != Info {
		t.Errorf("Logger level incorrect. Want: 1, Got: %d", level)
	}
}

func TestSetLogLevel(t *testing.T) {
	type test struct {
		level string
		want  Level
	}

	tests := []test{
		{want: Debug, level: "DEBUG"},
		{want: Info, level: "INFORMATION"},
		{want: Warn, level: "WARNING"},
		{want: Err, level: "ERROR"},
		{want: Fatal, level: "FATAL"},
		{want: Info, level: "INVALID"},
	}

	fl := NewFileLogger("test_log.log", time.RFC3339)

	for _, tc := range tests {
		fl.SetLevel(tc.level)
		if fl.level != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, fl.level)
		}
	}
}

func TestCheckLevel(t *testing.T) {
	type test struct {
		level Level
		want  bool
	}

	tests := []test{
		{level: Debug, want: true},
		{level: Info, want: true},
		{level: Warn, want: true},
		{level: Err, want: true},
		{level: Fatal, want: true},
		{level: 10, want: false},
	}

	fl := NewFileLogger("test_log.log", time.RFC3339)
	fl.SetLevel("DEBUG")

	for _, tc := range tests {
		got := fl.checkLevel(tc.level)
		if got != tc.want {
			t.Errorf("CheckLevel: Want: '%v', Got: '%v'", tc.want, got)
		}
	}
}

func TestLevelString(t *testing.T) {
	type test struct {
		level Level
		want  string
	}

	tests := []test{
		{level: Debug, want: "DEBUG"},
		{level: Info, want: "INFO "},
		{level: Warn, want: "WARN "},
		{level: Err, want: "ERROR"},
		{level: Fatal, want: "FATAL"},
	}

	for _, tc := range tests {
		got := tc.level.String()
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFormatMessage(t *testing.T) {
	str := "2021-07-25T00:00:00.000Z"
	testTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.Errorf("Could not parse string to time")
	}
	levelString := "DEBUG"
	testMsg := "Test message"
	msg := formatMessage(testTime, time.RFC3339, levelString, testMsg)
	want := "2021-07-25T00:00:00Z DEBUG: Test message\n"
	if strings.Compare(msg, want) != 0 {
		t.Errorf("FormatTime: Want: '%s', Got: '%s'", want, msg)
	}
}

func TestLog(t *testing.T) {
	var b bytes.Buffer
	want := "Test log message"
	if err := log(&b, want); err != nil {
		t.Fatalf("Could not log message")
	}
	got := b.String()
	if got != want {
		t.Errorf("log: Want: %s, Got: %s", want, got)
	}
}
