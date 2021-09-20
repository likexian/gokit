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

package xos

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
)

const (
	pidFile = "/tmp/testing.pid"
)

func TestValue(t *testing.T) {
	defer os.Remove(pidFile)

	p := Pid(pidFile)
	_, err := p.Value()
	assert.NotNil(t, err)

	err = xfile.WriteText(pidFile, "1")
	assert.Nil(t, err)

	pid, err := p.Value()
	assert.Nil(t, err)
	assert.Equal(t, pid, 1)
}

func TestAlive(t *testing.T) {
	defer os.Remove(pidFile)

	p := Pid(pidFile)
	_, err := p.Alive()
	assert.NotNil(t, err)

	err = xfile.WriteText(pidFile, "88888888")
	assert.Nil(t, err)

	_, err = p.Alive()
	assert.NotNil(t, err)

	err = xfile.WriteText(pidFile, "1")
	assert.Nil(t, err)

	pid, err := p.Alive()
	assert.Nil(t, err)
	assert.Equal(t, pid, 1)
}

func TestCreate(t *testing.T) {
	defer os.Remove(pidFile)

	err := xfile.WriteText(pidFile, "1")
	assert.Nil(t, err)

	p := Pid(pidFile)
	pid, err := p.Create()
	assert.NotNil(t, err)
	assert.Equal(t, pid, 1)

	err = xfile.WriteText(pidFile, "88888888")
	assert.Nil(t, err)

	_, err = p.Create()
	assert.Nil(t, err)
}

func TestConcurrency(t *testing.T) {
	defer os.Remove(pidFile)

	p := Pid(pidFile)
	for i := 0; i < 1000; i++ {
		go func() {
			_, err := p.Create()
			if err != nil {
				if !errors.Is(err, ErrPidLockFailed) && !errors.Is(err, ErrPidExists) {
					t.Errorf("Unexcepted error: %s", err)
				}
			}
		}()
	}

	time.Sleep(1 * time.Second)
}
