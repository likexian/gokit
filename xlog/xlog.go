/*
 * Copyright 2012-2019 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * A toolkit for Golang development
 * https://www.likexian.com/
 */

package xlog

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// LogLevel storing log level
type LogLevel int

// LogFile storing log file
type LogFile struct {
	Name          string
	Fd            *os.File
	Writer        io.Writer
	RotateType    string
	RotateNum     int64
	RotateSize    int64
	RotateNowDate string
	RotateNowSize int64
	RotateNextNum int64
}

// Logger storing logger
type Logger struct {
	LogFile  LogFile
	LogLevel LogLevel
	LogQueue chan string
	LogExit  chan bool
	Closed   bool
}

// OnceCache storing once cache
type OnceCache struct {
	Data map[string]int64
	sync.RWMutex
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
	"debug": DEBUG,
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
	"fatal": FATAL,
}

// log once cache
var onceCache = OnceCache{Data: map[string]int64{}}

// Version returns package version
func Version() string {
	return "0.2.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// New returns a new logger
func New(w io.Writer, level LogLevel) *Logger {
	return newLog(LogFile{Writer: w}, level)
}

// File returns a new file logger
func File(fname string, level LogLevel) (*Logger, error) {
	fd, err := openFile(fname)
	if err != nil {
		return nil, err
	}
	return newLog(LogFile{Name: fname, Writer: fd, Fd: fd}, level), nil
}

// openFile open file with flags
func openFile(fname string) (*os.File, error) {
	return os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

// newLogger returns a new file logger
func newLog(lf LogFile, level LogLevel) *Logger {
	l := &Logger{
		LogFile:  lf,
		LogLevel: level,
		LogQueue: make(chan string, 10000),
		LogExit:  make(chan bool),
		Closed:   false,
	}
	go l.writeLog()
	return l
}

// Close close the logger
func (l *Logger) Close() {
	l.Closed = true
	close(l.LogQueue)
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

// SetDailyRotate set daily log rotate
func (l *Logger) SetDailyRotate(rotateNum int64) error {
	return l.SetRotate("date", rotateNum, 0)
}

// SetSizeRotate set filesize log rotate
func (l *Logger) SetSizeRotate(rotateNum int64, rotateSize int64) error {
	return l.SetRotate("size", rotateNum, rotateSize)
}

// SetRotate set log rotate
// rotateType: date: daily rotate, size: filesize rotate
func (l *Logger) SetRotate(rotateType string, rotateNum int64, rotateSize int64) error {
	if l.LogFile.Name == "" {
		return errors.New("Only file log support rotate")
	}

	if rotateType != "date" && rotateType != "size" {
		return errors.New("Not support rotateType")
	}

	l.LogFile.RotateType = rotateType
	l.LogFile.RotateNum = rotateNum
	l.LogFile.RotateSize = rotateSize
	l.LogFile.RotateNowDate = time.Now().Format("2006-01-02")

	size, err := getFileSize(l.LogFile.Name)
	if err != nil {
		l.LogFile.RotateNowSize = 0
	} else {
		l.LogFile.RotateNowSize = size
	}

	if l.LogFile.RotateNum < 2 {
		return nil
	}

	list, err := getFileList(l.LogFile.Name)
	if err != nil {
		l.LogFile.RotateNextNum = 1
	} else {
		if int64(len(list)) < l.LogFile.RotateNum {
			l.LogFile.RotateNextNum = int64(len(list))
		} else {
			maxf := list[0]
			for _, v := range list {
				if v[0].(string) != l.LogFile.Name {
					if v[1].(int64) < maxf[1].(int64) {
						maxf = v
					}
				}
			}
			fs := strings.Split(maxf[0].(string), ".")
			num, _ := strconv.Atoi(fs[len(fs)-1])
			l.LogFile.RotateNextNum = int64(num)
		}
	}

	return nil
}

// writeLog get log from queue and write
func (l *Logger) writeLog() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			if l.LogFile.RotateType == "" {
				continue
			}
			if l.LogFile.RotateNum < 2 {
				continue
			}
			today := time.Now().Format("2006-01-02")
			if l.LogFile.RotateType == "date" {
				if today != l.LogFile.RotateNowDate {
					l.LogFile.RotateNowDate = today
					l.LogFile.RotateNowSize = 0
					l.rotateFile()
				}
			}
			if l.LogFile.RotateSize > 0 {
				if l.LogFile.RotateNowSize >= l.LogFile.RotateSize {
					l.LogFile.RotateNowDate = today
					l.LogFile.RotateNowSize = 0
					l.rotateFile()
				}
			}
		case s, ok := <-l.LogQueue:
			if !ok {
				l.LogExit <- true
				l.LogFile.Fd.Close()
				return
			}
			_, err := fmt.Fprintf(l.LogFile.Writer, s)
			if err == nil {
				l.LogFile.RotateNowSize += int64(len(s))
			}
		}
	}
}

// rotateFile do rotate log file
func (l *Logger) rotateFile() (err error) {
	l.LogFile.Fd.Close()

	err = os.Rename(l.LogFile.Name, fmt.Sprintf("%s.%d", l.LogFile.Name, l.LogFile.RotateNextNum))
	if err != nil {
		return
	}

	l.LogFile.RotateNextNum += 1
	if l.LogFile.RotateNextNum >= l.LogFile.RotateNum {
		l.LogFile.RotateNextNum = 1
	}

	fd, err := openFile(l.LogFile.Name)
	if err != nil {
		return err
	}

	l.LogFile.Fd = fd
	l.LogFile.Writer = fd

	return
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

	l.LogQueue <- fmt.Sprintf(str, args...)
}

// LogOnce do log a msg only one times
func (l *Logger) LogOnce(level string, msg string, args ...interface{}) {
	str := fmt.Sprintf("%s-%s", level, msg)
	key := md5Sum(fmt.Sprintf(str, args...))

	onceCache.RLock()
	_, ok := onceCache.Data[key]
	onceCache.RUnlock()
	if ok {
		return
	}

	onceCache.Lock()
	onceCache.Data[key] = time.Now().Unix()
	onceCache.Unlock()

	l.Log(level, msg, args...)
}

// Debug level msg logging
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Log("DEBUG", msg, args...)
}

// Info level msg logging
func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log("INFO", msg, args...)
}

// Warn level msg logging
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
	l.Close()
	l.exit(1)
}

// DebugOnce level msg logging
func (l *Logger) DebugOnce(msg string, args ...interface{}) {
	l.LogOnce("DEBUG", msg, args...)
}

// InfoOnce level msg logging
func (l *Logger) InfoOnce(msg string, args ...interface{}) {
	l.LogOnce("INFO", msg, args...)
}

// WarnOnce level msg logging
func (l *Logger) WarnOnce(msg string, args ...interface{}) {
	l.LogOnce("WARN", msg, args...)
}

// ErrorOnce level msg logging
func (l *Logger) ErrorOnce(msg string, args ...interface{}) {
	l.LogOnce("ERROR", msg, args...)
}

// FatalOnce level msg logging, followed by a call to os.Exit(1)
func (l *Logger) FatalOnce(msg string, args ...interface{}) {
	l.LogOnce("FATAL", msg, args...)
}

// exit wait for queue empty and call os.Exit()
func (l *Logger) exit(code int) {
	select {
	case <-l.LogExit:
		os.Exit(code)
	}
}

// getFileSize returns file size
func getFileSize(fname string) (int64, error) {
	f, err := os.Stat(fname)
	if err != nil {
		return 0, err
	}

	return f.Size(), nil
}

// getFileList returns file list
func getFileList(fname string) (result [][]interface{}, err error) {
	result = [][]interface{}{}

	fs, err := filepath.Glob(fname + "*")
	if err != nil {
		return
	}

	for _, f := range fs {
		fd, e := os.Stat(f)
		if e != nil {
			err = e
			return
		}
		result = append(result, []interface{}{f, fd.ModTime().Unix()})
	}

	return
}

// md5Sum returns hex md5 of string
func md5Sum(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
