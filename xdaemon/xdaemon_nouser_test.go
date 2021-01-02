/*
 * Copyright 2012-2021 Li Kexian
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

package xdaemon

import (
	"os"
	"os/exec"
	"testing"

	"github.com/likexian/gokit/assert"
)

func testNoUser(t *testing.T) {
	c := Config{
		Pid:   "",
		Log:   "",
		User:  "",
		Chdir: "",
	}

	err := c.Daemon()
	assert.Nil(t, err)
}

func TestNoUser(t *testing.T) {
	if os.Getenv("TestCase") != "" && os.Getenv("TestCase") != "TestNoUser" {
		return
	}

	if os.Getenv("TestCase") != "" {
		testNoUser(t)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestNoUser")
	cmd.Env = append(os.Environ(), "TestCase=TestNoUser")
	err := cmd.Run()
	if err == nil {
		return
	}

	t.Errorf("Test expect to be daemon")
}
