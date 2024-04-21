/*
 * Copyright 2012-2024 Li Kexian
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
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/likexian/gokit/xcache"
	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xhash"
)

// Log level const
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Log prefix flag, similar to golang log package
const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime
)

// log level mapper
var levelMap = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

// log once cache
var onceCache = xcache.New(xcache.MemoryCache)

// LogLevel storing log level
type LogLevel int

// LogFlag storing log flag
type LogFlag int

// Logger storing logger
type Logger struct {
	logFile   logFile
	logLevel  LogLevel
	logFlag   LogFlag
	logQueue  chan string
	logExit   chan bool
	logClosed bool
	sync.RWMutex
}

// logFile storing log file info
type logFile struct {
	name          string
	fd            *os.File
	writer        io.Writer
	rotateType    string
	rotateNum     int64
	rotateSize    int64
	rotateNowDate string
	rotateNowSize int64
	rotateNextNum int64
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
	return "Licensed under the Apache License 2.0"
}

// New returns a new logger
func New(w io.Writer, level LogLevel) *Logger {
	return newLog(logFile{writer: w}, level, LstdFlags)
}

// File returns a new file logger
func File(fname string, level LogLevel) (*Logger, error) {
	fd, err := openFile(fname)
	if err != nil {
		return nil, err
	}
	return newLog(logFile{name: fname, writer: fd, fd: fd}, level, LstdFlags), nil
}

// openFile open file with flags
func openFile(fname string) (*os.File, error) {
	return os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

// newLogger returns a new file logger
func newLog(lf logFile, level LogLevel, flag LogFlag) *Logger {
	l := &Logger{
		logFile:   lf,
		logLevel:  level,
		logFlag:   flag,
		logQueue:  make(chan string, 10000),
		logExit:   make(chan bool),
		logClosed: false,
	}
	go l.writeLog()
	return l
}

// Close close the logger
func (l *Logger) Close() {
	l.Lock()
	l.logClosed = true
	close(l.logQueue)
	l.Unlock()
	<-l.logExit
}

// SetLevel set the log level by int level
func (l *Logger) SetLevel(level LogLevel) {
	l.Lock()
	l.logLevel = level
	l.Unlock()
}

// SetFlag set the log flag
func (l *Logger) SetFlag(flag LogFlag) {
	l.Lock()
	l.logFlag = flag
	l.Unlock()
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
	l.Lock()
	defer l.Unlock()

	if l.logFile.name == "" {
		return errors.New("xlog: rotate require log to file")
	}

	if rotateType != "date" && rotateType != "size" {
		return fmt.Errorf("xlog: not supported rotate type: %s", rotateType)
	}

	l.logFile.rotateType = rotateType
	l.logFile.rotateNum = rotateNum
	l.logFile.rotateSize = rotateSize
	l.logFile.rotateNowDate = time.Now().Format("2006-01-02")

	size, err := xfile.Size(l.logFile.name)
	if err != nil {
		l.logFile.rotateNowSize = 0
	} else {
		l.logFile.rotateNowSize = size
	}

	if l.logFile.rotateNum < 2 {
		return nil
	}

	list, err := getFileList(l.logFile.name)
	if err != nil {
		l.logFile.rotateNextNum = 1
	} else {
		if int64(len(list)) < l.logFile.rotateNum {
			l.logFile.rotateNextNum = int64(len(list))
		} else {
			maxf := list[0]
			for _, v := range list {
				if v[0].(string) != l.logFile.name {
					if v[1].(int64) < maxf[1].(int64) {
						maxf = v
					}
				}
			}
			fs := strings.Split(maxf[0].(string), ".")
			num, _ := strconv.Atoi(fs[len(fs)-1])
			l.logFile.rotateNextNum = int64(num)
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
			l.RLock()
			lf := l.logFile
			l.RUnlock()
			if lf.rotateType == "" {
				continue
			}
			if lf.rotateNum < 2 {
				continue
			}
			today := time.Now().Format("2006-01-02")
			if lf.rotateType == "date" {
				if today != lf.rotateNowDate {
					l.Lock()
					l.logFile.rotateNowDate = today
					l.logFile.rotateNowSize = 0
					l.Unlock()
					_ = l.rotateFile()
				}
			}
			if lf.rotateSize > 0 {
				if lf.rotateNowSize >= lf.rotateSize {
					l.Lock()
					l.logFile.rotateNowDate = today
					l.logFile.rotateNowSize = 0
					l.Unlock()
					_ = l.rotateFile()
				}
			}
		case s, ok := <-l.logQueue:
			l.Lock()
			if !ok {
				l.logExit <- true
				l.logFile.fd.Close()
				l.Unlock()
				return
			}
			_, err := fmt.Fprint(l.logFile.writer, s)
			if err == nil {
				l.logFile.rotateNowSize += int64(len(s))
			}
			l.Unlock()
		}
	}
}

// rotateFile do rotate log file
func (l *Logger) rotateFile() (err error) {
	l.Lock()
	defer l.Unlock()

	l.logFile.fd.Close()

	err = os.Rename(l.logFile.name, fmt.Sprintf("%s.%d", l.logFile.name, l.logFile.rotateNextNum))
	if err != nil {
		return
	}

	l.logFile.rotateNextNum++
	if l.logFile.rotateNextNum >= l.logFile.rotateNum {
		l.logFile.rotateNextNum = 1
	}

	fd, err := openFile(l.logFile.name)
	if err != nil {
		return err
	}

	l.logFile.fd = fd
	l.logFile.writer = fd

	return
}

// Log do log a msg
func (l *Logger) Log(level LogLevel, msg string, args ...interface{}) {
	if l.logClosed {
		return
	}

	if l.logLevel > level {
		return
	}

	if _, ok := levelMap[level]; !ok {
		return
	}

	logTime := ""
	now := time.Now()
	if l.logFlag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if l.logFlag&LUTC != 0 {
			now = now.UTC()
		}
		if l.logFlag&Ldate != 0 {
			logTime += fmt.Sprintf("%s ", now.Format("2006-01-02"))
		}
		if l.logFlag&Ltime != 0 {
			logTime += fmt.Sprintf("%s ", now.Format("15:04:05"))
		}
		if l.logFlag&Lmicroseconds != 0 {
			logTime = fmt.Sprintf("%s.%d ", strings.TrimSpace(logTime), now.Nanosecond()/1e3)
		}
	}

	logFile := ""
	if l.logFlag&(Llongfile|Lshortfile) != 0 {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			logFile = "???:? "
		} else {
			logFile = fmt.Sprintf("%s:%d ", file, line)
		}
		if l.logFlag&Lshortfile != 0 {
			ls := strings.Split(logFile, "/")
			logFile = ls[len(ls)-1]
		}
	}

	str := fmt.Sprintf("%s%s[%s] %s\n", logTime, logFile, levelMap[level], msg)
	l.logQueue <- fmt.Sprintf(str, args...)
}

// LogOnce do log a msg only one times within one hour
func (l *Logger) LogOnce(level LogLevel, msg string, args ...interface{}) {
	str := fmt.Sprintf("%d-%s", level, msg)
	key := xhash.Md5(fmt.Sprintf(str, args...)).Hex()

	v := onceCache.Get(key)
	if v != nil {
		return
	}

	_ = onceCache.Set(key, 1, 3600)
	l.Log(level, msg, args...)
}

// Debug level msg logging
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Log(DEBUG, msg, args...)
}

// Info level msg logging
func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log(INFO, msg, args...)
}

// Warn level msg logging
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Log(WARN, msg, args...)
}

// Error level msg logging
func (l *Logger) Error(msg string, args ...interface{}) {
	l.Log(ERROR, msg, args...)
}

// Fatal level msg logging, followed by a call to os.Exit(1)
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Log(FATAL, msg, args...)
	l.Close()
	os.Exit(1)
}

// DebugOnce level msg logging
func (l *Logger) DebugOnce(msg string, args ...interface{}) {
	l.LogOnce(DEBUG, msg, args...)
}

// InfoOnce level msg logging
func (l *Logger) InfoOnce(msg string, args ...interface{}) {
	l.LogOnce(INFO, msg, args...)
}

// WarnOnce level msg logging
func (l *Logger) WarnOnce(msg string, args ...interface{}) {
	l.LogOnce(WARN, msg, args...)
}

// ErrorOnce level msg logging
func (l *Logger) ErrorOnce(msg string, args ...interface{}) {
	l.LogOnce(ERROR, msg, args...)
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
