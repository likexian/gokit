/*
 * Copyright 2012-2020 Li Kexian
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
		output := strings.TrimSpace(stderr.String())
		if !strings.Contains(output, LogLine) {
			t.Errorf("Test got %s, expect %s", output, LogLine)
		}
		return
	}

	t.Errorf("Test got err %s, expect exit status 1", err.Error())
}
