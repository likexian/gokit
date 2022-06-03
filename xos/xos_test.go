/*
 * Copyright 2012-2022 Li Kexian
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
	"os"
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestGetenv(t *testing.T) {
	assert.Equal(t, Getenv("TEST", ""), "")

	os.Setenv("TEST", "gokit")
	assert.Equal(t, Getenv("TEST", ""), "gokit")
}

func TestExec(t *testing.T) {
	_, _, err := Exec("xx")
	assert.NotNil(t, err)

	stdout, stderr, err := Exec("ls", "-lh")
	assert.Nil(t, err)
	assert.NotEqual(t, stdout, "")
	assert.Equal(t, stderr, "")
}

func TestTimeoutExec(t *testing.T) {
	_, _, err := TimeoutExec(1, "xxx")
	assert.NotNil(t, err)

	_, _, err = TimeoutExec(1, "sleep", "3")
	assert.NotNil(t, err)

	stdout, stderr, err := TimeoutExec(3, "sleep", "1")
	assert.Nil(t, err)
	assert.Equal(t, stdout, "")
	assert.Equal(t, stderr, "")
}

func TestLookupUser(t *testing.T) {
	uid, gid, err := LookupUser("nobody")
	assert.Nil(t, err)
	assert.True(t, uid > 0)
	assert.True(t, gid > 0)
}

func TestGetPwd(t *testing.T) {
	pwd := GetPwd()
	assert.NotEqual(t, pwd, "", "pwd expect to be not empty")

	pwd = GetProcPwd()
	assert.NotEqual(t, pwd, "", "pwd expect to be not empty")
}

func TestSetid(t *testing.T) {
	err := SetUID(0)
	assert.Nil(t, err)

	err = SetGID(0)
	assert.Nil(t, err)

	err = SetUser("root")
	assert.Nil(t, err)

	err = SetUser("xxx")
	assert.NotNil(t, err)

	/*
		err = SetUser("nobody")
		assert.Nil(t, err)

		err = SetUID(0)
		assert.NotNil(t, err)

		err = SetGID(0)
		assert.NotNil(t, err)

		err = SetUser("root")
		assert.NotNil(t, err)
	*/
}
