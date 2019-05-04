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
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const LogLine = "This line will show before exit"

func LogFatal() {
	log := New(os.Stderr, DEBUG)
	log.Fatal(LogLine)
}

func TestFatal(t *testing.T) {
	if os.Getenv("TestFatal") == "1" {
		LogFatal()
		return
	}

	var stderr bytes.Buffer
	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "TestFatal=1")
	cmd.Stderr = &stderr

	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		output := strings.TrimSpace(string(stderr.Bytes()))
		if !strings.Contains(output, LogLine) {
			t.Errorf("Test got %s, expect %s", output, LogLine)
		}
		return
	}

	t.Errorf("Test got err %s, expect exit status 1", err.Error())
}
