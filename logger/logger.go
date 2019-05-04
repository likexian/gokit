/*
 * Go module for do logging
 * https://www.likexian.com/
 *
 * Copyright 2015, Kexian Li
 * Released under the Apache License, Version 2.0
 *
 */


package logger


import (
    "fmt"
    "strings"
    "time"
    "sync"
    "io"
    "os"
)


type Level int

type Logger struct {
    Writer io.Writer
    Level Level
    sync.Mutex
}


const (
    DEBUG Level = iota
    INFO
    NOTICE
    WARNING
    ERROR
    CRITICAL
)


var levels = map[string]Level {
    "debug": DEBUG,
    "info": INFO,
    "notice": NOTICE,
    "warning": WARNING,
    "error": ERROR,
    "critical": CRITICAL,
}


func Version() string {
    return "0.1.0"
}


func Author() string {
    return "[Li Kexian](https://www.likexian.com/)"
}


func License() string {
    return "Apache License, Version 2.0"
}


func New(w io.Writer, level Level) *Logger {
    l := &Logger{Writer: w, Level: level}
    return l
}


func File(fname string, level Level) (*Logger, error) {
    fd, err := os.OpenFile(fname, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
    if err != nil {
        return nil, err
    }
    return New(fd, level), nil
}


func (l *Logger) SetLevel(level Level) {
    l.Lock()
    defer l.Unlock()
    l.Level = level
}


func (l *Logger) SetLevelString(level string) error {
    value := l.GetLevelByString(level)
    if value >= 0 {
        l.SetLevel(value)
    }
    return fmt.Errorf("%s is invalid level", level)
}


func (l *Logger) GetLevelByString(level string) Level {
    level = strings.ToLower(level)
    if value, ok := levels[level]; ok {
        return value
    }
    return -1
}


func (l *Logger) Log(level string, msg string, args ...interface{}) error {
    l.Lock()
    defer l.Unlock()

    value := l.GetLevelByString(level)
    if l.Level > value {
        return nil
    }

    now := time.Now().Format("2006-01-02 15:04:05")
    str := fmt.Sprintf("%s [%s] %s\n", now, strings.ToUpper(level), msg)
    _, err := fmt.Fprintf(l.Writer, str, args...)

    return err
}


func (l *Logger) Debug(msg string, args ...interface{}) error {
    return l.Log("DEBUG", msg, args...)
}


func (l *Logger) Info(msg string, args ...interface{}) error {
    return l.Log("INFO", msg, args...)
}


func (l *Logger) Notice(msg string, args ...interface{}) error {
    return l.Log("NOTICE", msg, args...)
}


func (l *Logger) Warning(msg string, args ...interface{}) error {
    return l.Log("WARNING", msg, args...)
}


func (l *Logger) Error(msg string, args ...interface{}) error {
    return l.Log("ERROR", msg, args...)
}


func (l *Logger) Critical(msg string, args ...interface{}) error {
    return l.Log("CRITICAL", msg, args...)
}
