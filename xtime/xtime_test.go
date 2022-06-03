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

package xtime

import (
	"errors"
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestGetSecond(t *testing.T) {
	tm := Now()

	s := String()
	assert.NotEqual(t, s, "")

	n := S()
	assert.True(t, n >= tm.Unix())

	n = Ns()
	assert.True(t, n >= tm.UnixNano())

	n = Us()
	assert.True(t, n >= tm.UnixNano()/int64(time.Microsecond))

	n = Ms()
	assert.True(t, n >= tm.UnixNano()/int64(time.Millisecond))
}

func TestSleep(t *testing.T) {
	tm := S()
	Sleep(1)
	assert.Equal(t, S(), tm+1)
}

func TestUsleep(t *testing.T) {
	tm := Us()
	Usleep(1000)
	assert.True(t, Us() >= tm+1000)
}

func TestStrToTime(t *testing.T) {
	n, err := StrToTime("2019-03-11")
	assert.Nil(t, err)
	en, _ := time.ParseInLocation("2006-01-02", "2019-03-11", time.Local)
	assert.Equal(t, n, en.Unix())

	n, err = StrToTime("2019-03-11 22:23:24")
	assert.Nil(t, err)
	en, _ = time.ParseInLocation("2006-01-02 15:04:05", "2019-03-11 22:23:24", time.Local)
	assert.Equal(t, n, en.Unix())

	n, err = StrToTime("2019-03-11T22:23:24Z", "2006-01-02T15:04:05Z")
	assert.Nil(t, err)
	en, _ = time.ParseInLocation("2006-01-02T15:04:05Z", "2019-03-11T22:23:24Z", time.Local)
	assert.Equal(t, n, en.Unix())

	n, err = StrToTime("Mon, 11 Mar 2019 22:23:24", "Mon, 02 Jan 2006 15:04:05")
	assert.Nil(t, err)
	en, _ = time.ParseInLocation("Mon, 02 Jan 2006 15:04:05", "Mon, 11 Mar 2019 22:23:24", time.Local)
	assert.Equal(t, n, en.Unix())

	n, err = StrToTime("2019-03-11T22:23:24Z", time.RFC3339)
	assert.Nil(t, err)
	en, _ = time.Parse(time.RFC3339, "2019-03-11T22:23:24Z")
	assert.Equal(t, n, en.Unix())

	n, err = StrToTime("2019-03-11T22:23:24+08:00", time.RFC3339)
	assert.Nil(t, err)
	en, _ = time.Parse(time.RFC3339, "2019-03-11T22:23:24+08:00")
	assert.Equal(t, n, en.Unix())

	_, err = StrToTime("2019-03-11T22:23:24Z")
	assert.NotNil(t, err)
}

func TestTimeToStr(t *testing.T) {
	s := TimeToStr(0)
	assert.Equal(t, s, time.Unix(0, 0).Format("2006-01-02 15:04:05"))

	s = TimeToStr(1552233600)
	assert.Equal(t, s, time.Unix(1552233600, 0).Format("2006-01-02 15:04:05"))

	s = TimeToStr(time.Now().Unix())
	assert.Equal(t, s, time.Now().Format("2006-01-02 15:04:05"))

	s = TimeToStr(1552314204, "2006-01-02T15:04:05Z")
	assert.Equal(t, s, time.Unix(1552314204, 0).Format("2006-01-02T15:04:05Z"))

	s = TimeToStr(1552314204, "Mon, 02 Jan 2006 15:04:05")
	assert.Equal(t, s, time.Unix(1552314204, 0).Format("Mon, 02 Jan 2006 15:04:05"))
}

func TestWithTimeout(t *testing.T) {
	r, _ := WithTimeout(func() interface{} { return 10000 }, 1*time.Second)
	assert.Equal(t, <-r, 10000)

	r, _ = WithTimeout(func() interface{} { return errors.New("some error") }, 1*time.Second)
	assert.NotNil(t, <-r)

	r, _ = WithTimeout(func() interface{} { Sleep(2); t.Log("i run after timeout"); return 10000 }, 1*time.Second)
	assert.Equal(t, <-r, ErrTimeouted)

	r, cancel := WithTimeout(func() interface{} { Sleep(2); t.Log("i run after cancel"); return 10000 }, 1*time.Second)
	cancel()
	assert.Equal(t, <-r, ErrCanceled)

	Sleep(3)
}

func TestSetTimeout(t *testing.T) {
	r, _ := SetTimeout(func() interface{} { return 10000 }, 1*time.Second)
	start := S()
	assert.Equal(t, <-r, 10000)
	assert.Equal(t, S()-start, int64(1))

	r, _ = SetTimeout(func() interface{} { return errors.New("some error") }, 1*time.Second)
	start = S()
	assert.NotNil(t, <-r)
	assert.Equal(t, S()-start, int64(1))

	r, cancel := SetTimeout(func() interface{} { Sleep(2); t.Log("i will not run"); return 10000 }, 1*time.Second)
	start = S()
	cancel()
	assert.Equal(t, <-r, ErrCanceled)
	assert.Equal(t, S()-start, int64(0))

	Sleep(3)
}

func TestSetInterval(t *testing.T) {
	r, _ := SetInterval(func() interface{} { return 10000 }, 1*time.Second)
	start := S()
	assert.Equal(t, <-r, 10000)
	assert.Equal(t, S()-start, int64(1))

	r, _ = SetInterval(func() interface{} { return errors.New("some error") }, 1*time.Second)
	start = S()
	assert.NotNil(t, <-r)
	assert.Equal(t, S()-start, int64(1))

	r, cancel := SetInterval(func() interface{} { Sleep(2); t.Log("i will not run"); return 10000 }, 1*time.Second)
	start = S()
	cancel()
	assert.Equal(t, <-r, ErrCanceled)
	assert.Equal(t, S()-start, int64(0))

	r, cancel = SetInterval(func() interface{} { return 10000 }, 1*time.Second)
	start = S()
	ints := []int{}
	for i := 0; i < 3; i++ {
		ints = append(ints, (<-r).(int))
	}
	cancel()
	assert.Equal(t, <-r, ErrCanceled)
	assert.Equal(t, ints, []int{10000, 10000, 10000})
	assert.Equal(t, S()-start, int64(3))

	Sleep(3)
}
