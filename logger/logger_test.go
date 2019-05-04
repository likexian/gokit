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
	"time"
	"testing"
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
