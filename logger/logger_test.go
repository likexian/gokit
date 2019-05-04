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
	"testing"
)

func TestLogger(t *testing.T) {
	log := New(os.Stderr, DEBUG)
	log.Debug("This is Debug")
	log.Info("This is Info")
	log.Warn("This is Warn")
	log.Error("This is Error")

	log.Info("Now setting level to Info")
	log.SetLevel(INFO)
	log.Debug("This is Debug")
	log.Info("This is Info")

	log.SetLevel(INFO)
	log.Info("Now setting level to Error")
	log.SetLevelString("ERROR")
	log.Warn("This is Warn")
	log.Error("This is Error")

	flog, err := File("test.log", DEBUG)
	if err != nil {
		panic(err)
	}
	flog.Debug("This is Debug")
	flog.Info("This is Info")
	flog.Warn("This is Warn")
	flog.Error("This is Error")
}
