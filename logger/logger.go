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

// LogLevel storing log level
type LogLevel int

// LogFile storing log file
type LogFile struct {
	Name 		string
	Fd          *os.File
	Writer 		io.Writer
}

// Logger storing logger
type Logger struct {
	LogFile 	LogFile
	LogLevel  	LogLevel
	Queue  		chan string
	Closed 		bool
	sync.Mutex
}

// Log level const
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// log level mapper
var levels = map[string]LogLevel{
	"debug":    DEBUG,
	"info":     INFO,
	"warn":  	WARN,
	"error":    ERROR,
	"fatal": 	FATAL,
}

// Version returns package version
func Version() string {
	return "0.7.0"
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
func New(w io.Writer, level LogLevel) *Logger {
	return newFile(LogFile{Writer: w}, level)
}

// File returns a new file logger
func File(fname string, level LogLevel) (*Logger, error) {
	fd, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return newFile(LogFile{fname, fd, fd}, level), nil
}

// newLogger returns a new file logger
func newFile(lf LogFile, level LogLevel) *Logger {
	l := &Logger{
		LogFile: 	lf,
		LogLevel: 	level,
		Queue: 		make(chan string, 10000),
		Closed: 	false,
	}
	go l.writeLog()
	return l
}

// Close close the logger
func (l *Logger) Close() {
	l.Closed = true
	close(l.Queue)
}

// SetLevel set the log level by int level
func (l *Logger) SetLevel(level LogLevel) {
	l.LogLevel = level
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
func (l *Logger) GetLevelByString(level string) LogLevel {
	level = strings.ToLower(level)
	if value, ok := levels[level]; ok {
		return value
	}
	return -1
}

// writeLog get log from queue and write
func (l *Logger) writeLog() {
	for {
		t, ok := <- l.Queue
		if !ok {
			l.LogFile.Fd.Close()
			return
		}
		_, err := fmt.Fprintf(l.LogFile.Writer, t)
		if err != nil {
		}
	}
}

// Log do log a msg
func (l *Logger) Log(level string, msg string, args ...interface{}) {
	if l.Closed {
		return
	}

	value := l.GetLevelByString(level)
	if l.LogLevel > value {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [%s] %s\n", now, strings.ToUpper(level), msg)

	l.Queue <- fmt.Sprintf(str, args...)
}

// Debug level msg logging
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Log("DEBUG", msg, args...)
}

// Info level msg logging
func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log("INFO", msg, args...)
}

// Warning level msg logging
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Log("WARN", msg, args...)
}

// Error level msg logging
func (l *Logger) Error(msg string, args ...interface{}) {
	l.Log("ERROR", msg, args...)
}

// Fatal level msg logging, followed by a call to os.Exit(1)
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Log("FATAL", msg, args...)
	os.Exit(1)
}
