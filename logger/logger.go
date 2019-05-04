/*
 * Go module for doing logging
 * https://www.likexian.com/
 *
 * Copyright 2015-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Level storing log level
type Level int

// Logger storing logger
type Logger struct {
	Writer io.Writer
	Level  Level
	sync.Mutex
}

// Log level const
const (
	DEBUG Level = iota
	INFO
	NOTICE
	WARNING
	ERROR
	CRITICAL
)

// log level mapper
var levels = map[string]Level{
	"debug":    DEBUG,
	"info":     INFO,
	"notice":   NOTICE,
	"warning":  WARNING,
	"error":    ERROR,
	"critical": CRITICAL,
}

// Version returns package version
func Version() string {
	return "0.3.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// New returns a new logger
func New(w io.Writer, level Level) *Logger {
	l := &Logger{Writer: w, Level: level}
	return l
}

// File returns a new file logger
func File(fname string, level Level) (*Logger, error) {
	fd, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return New(fd, level), nil
}

// SetLevel set the log level by int level
func (l *Logger) SetLevel(level Level) {
	l.Lock()
	defer l.Unlock()
	l.Level = level
}

// SetLevelString set the log level by string level
func (l *Logger) SetLevelString(level string) error {
	value := l.GetLevelByString(level)
	if value >= 0 {
		l.SetLevel(value)
	}
	return fmt.Errorf("%s is invalid level", level)
}

// GetLevelByString returns log level by string level
func (l *Logger) GetLevelByString(level string) Level {
	level = strings.ToLower(level)
	if value, ok := levels[level]; ok {
		return value
	}
	return -1
}

// Log do log a msg
func (l *Logger) Log(level string, msg string, args ...interface{}) error {
	value := l.GetLevelByString(level)
	if l.Level > value {
		return nil
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [%s] %s\n", now, strings.ToUpper(level), msg)

	l.Lock()
	_, err := fmt.Fprintf(l.Writer, str, args...)
	l.Unlock()

	return err
}

// Debug level msg logging
func (l *Logger) Debug(msg string, args ...interface{}) error {
	return l.Log("DEBUG", msg, args...)
}

// Info level msg logging
func (l *Logger) Info(msg string, args ...interface{}) error {
	return l.Log("INFO", msg, args...)
}

// Notice level msg logging
func (l *Logger) Notice(msg string, args ...interface{}) error {
	return l.Log("NOTICE", msg, args...)
}

// Warning level msg logging
func (l *Logger) Warning(msg string, args ...interface{}) error {
	return l.Log("WARNING", msg, args...)
}

// Error level msg logging
func (l *Logger) Error(msg string, args ...interface{}) error {
	return l.Log("ERROR", msg, args...)
}

// Critical level msg logging
func (l *Logger) Critical(msg string, args ...interface{}) error {
	return l.Log("CRITICAL", msg, args...)
}
