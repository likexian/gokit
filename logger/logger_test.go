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
	"os"
	"sync"
	"testing"
	"time"
)

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
	log.Error("Test logger closed")

	// log to file
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

func TestConcurrency(t *testing.T) {
	// log to stderr
	log := New(os.Stderr, DEBUG)
	for i := 0; i < 10; i++ {
		go func(i int) {
			log.Info("This is %d", i)
		}(i)
	}

	// log to file
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
	log, err := File("rotate.log", DEBUG)
	if err != nil {
		panic(err)
	}

	// set rotate by filesize
	log.SetSizeRotate(10, 100000)
	for i := 0; i < 200000; i++ {
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
