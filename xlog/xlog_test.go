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
	"github.com/likexian/gokit/assert"
	"os"
	"sync"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestLogger(t *testing.T) {
	// log to stderr
	log := New(os.Stderr, DEBUG)
	log.Info("Now setting level to Debug")
	log.Debug("This is Debug")
	log.Info("This is Info")
	log.Warn("This is Warn")
	log.Error("This is Error")
	log.Error("This is %s", "Args")
	log.Error("")

	// test SetLevel
	log.Info("Now setting level to Info")
	log.SetLevel(INFO)
	log.Debug("This is Debug shall NOT! shown")
	log.Info("This is Info")
	log.Warn("This is Warn")
	log.Error("This is Error")
	log.Error("This is %s", "Args")
	log.Error("")

	// test SetLevelString
	log.Info("Now setting level to Error")
	log.SetLevelString("ERROR")
	log.Debug("This is Debug shall NOT! shown")
	log.Info("This is Info shall NOT! shown")
	log.Warn("This is Warn shall NOT! shown")
	log.Error("This is Error")
	log.Error("This is %s", "Args")
	log.Error("")
	log.Close()
	log.Error("Test log closed")

	// log to file
	defer os.Remove("test.log")
	log, err := File("test.log", DEBUG)
	if err != nil {
		panic(err)
	}
	log.Debug("This is Debug")
	log.Info("This is Info")
	log.Warn("This is Warn")
	log.Error("This is Error")
	log.Error("This is %s", "Args")
	log.Close()

	// wait for queue empty
	time.Sleep(1 * time.Second)
}

func TestLogOnce(t *testing.T) {
	// log to stderr
	log := New(os.Stderr, DEBUG)

	log.WarnOnce("This only log once")
	log.WarnOnce("This only log once")
	log.WarnOnce("This only log once, %d", 1)
	log.WarnOnce("This only log once, %d", 1)
	log.WarnOnce("This only log once, %d", 2)

	// wait for queue empty
	time.Sleep(1 * time.Second)
}

func TestConcurrency(t *testing.T) {
	// log to stderr
	log := New(os.Stderr, DEBUG)
	for i := 0; i < 10; i++ {
		go func(i int) {
			log.Info("This is %d", i)
		}(i)
	}

	// log to file
	defer os.Remove("test.log")
	flog, err := File("test.log", DEBUG)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			flog.Info("This is %d", i)
		}(i)
	}

	// wait for queue empty
	time.Sleep(1 * time.Second)
	log.Close()
	flog.Close()
}

func TestLogRotate(t *testing.T) {
	var wg sync.WaitGroup

	// log to file
	defer os.Remove("test.log")
	defer os.Remove("test.log.1")
	defer os.Remove("test.log.2")
	log, err := File("test.log", DEBUG)
	if err != nil {
		panic(err)
	}

	// set rotate by filesize
	log.SetSizeRotate(3, 100000)
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			log.Info("This is a log line of log file by log thread: %d", i)
			wg.Done()
		}(i)
	}

	// wait for log end
	time.Sleep(3 * time.Second)
	wg.Wait()
	log.Close()
}