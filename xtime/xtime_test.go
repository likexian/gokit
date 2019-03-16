/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xtime

import (
	"errors"
	"github.com/likexian/gokit/assert"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
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
	assert.Equal(t, n, int64(1552233600))

	n, err = StrToTime("2019-03-11 22:23:24")
	assert.Nil(t, err)
	assert.Equal(t, n, int64(1552314204))

	n, err = StrToTime("2019-03-11T22:23:24", "2006-01-02T15:04:05")
	assert.Nil(t, err)
	assert.Equal(t, n, int64(1552314204))

	n, err = StrToTime("Mon, 11 Mar 2019 22:23:24", "Mon, 02 Jan 2006 15:04:05")
	assert.Nil(t, err)
	assert.Equal(t, n, int64(1552314204))

	n, err = StrToTime("2019-03-11T22:23:24")
	assert.NotNil(t, err)
}

func TestTimeToStr(t *testing.T) {
	s := TimeToStr(0)
	assert.Equal(t, s, "1970-01-01 08:00:00")

	s = TimeToStr(1552233600)
	assert.Equal(t, s, "2019-03-11 00:00:00")

	s = TimeToStr(1552314204)
	assert.Equal(t, s, "2019-03-11 22:23:24")

	s = TimeToStr(1552314204, "2006-01-02T15:04:05")
	assert.Equal(t, s, "2019-03-11T22:23:24")

	s = TimeToStr(1552314204, "Mon, 02 Jan 2006 15:04:05")
	assert.Equal(t, s, "Mon, 11 Mar 2019 22:23:24")
}

func TestWithTimeout(t *testing.T) {
	n, err := WithTimeout(func() interface{} { return 10000 }, 1*time.Second)
	assert.Nil(t, err)
	assert.Equal(t, n, 10000)

	n, err = WithTimeout(func() interface{} { return errors.New("some error") }, 1*time.Second)
	assert.Nil(t, n)
	assert.NotNil(t, err)
	assert.NotNil(t, err, ErrTimeout)

	n, err = WithTimeout(func() interface{} { Sleep(2); return 10000 }, 1*time.Second)
	assert.Equal(t, err, ErrTimeout)
	assert.NotEqual(t, n, 10000)
}
