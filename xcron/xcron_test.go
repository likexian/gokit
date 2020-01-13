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

package xcron

import (
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xtime"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestAddSet(t *testing.T) {
	c := New()

	_, err := c.Add("@error", func() {})
	assert.NotNil(t, err)

	id, err := c.Add("@every second", func() { t.Log("add a echo") })
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	assert.Equal(t, c.Len(), 1)

	time.Sleep(1 * time.Second)
	err = c.Set(id, "@every second", func() { t.Log("set a echo") })
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	assert.Equal(t, c.Len(), 1)

	time.AfterFunc(3*time.Second, func() { c.Del(id) })
	c.Wait()
	assert.Equal(t, c.Len(), 0)
}

func TestEmpty(t *testing.T) {
	c := New()

	for i := 0; i < 3; i++ {
		id, err := c.Add("@every second", func() { t.Log("add 3 echo") })
		assert.Nil(t, err)
		assert.NotEqual(t, id, "")
	}

	assert.Equal(t, c.Len(), 3)
	time.AfterFunc(3*time.Second, func() { c.Empty() })
	c.Wait()
	assert.Equal(t, c.Len(), 0)
}

func TestAddEchoChannel(t *testing.T) {
	r := make(chan interface{}, 10)
	echo := func(i int) int {
		return i
	}

	c := New()

	id, err := c.Add("@every 3 second", func() { r <- echo(int(xtime.S())) }, func() { close(r) })
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")

	time.AfterFunc(3*time.Second, func() { c.Del(id) })

	d := []interface{}{}
	for {
		i, ok := <-r
		if !ok {
			break
		}
		d = append(d, i)
	}
	assert.Len(t, d, 1)
}

func TestParse(t *testing.T) {
	tests := []struct {
		in  string
		out Rule
		err error
	}{
		{"", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"*", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * * * *", Rule{[]int{0}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * * * * *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * * * jan *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{1}, []int{}}, nil},
		{"* * * * jan-mar *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{1, 2, 3}, []int{}}, nil},
		{"* * * * jan,feb,mar *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{1, 2, 3}, []int{}}, nil},
		{"* * * * * sun", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{0}}, nil},
		{"* * * * * sun-tue", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{0, 1, 2}}, nil},
		{"* * * * * sun,mon,tue", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{0, 1, 2}}, nil},

		{"1 * * * * *", Rule{[]int{1}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* 1 * * * *", Rule{[]int{}, []int{1}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * 1 * * *", Rule{[]int{}, []int{}, []int{1}, []int{}, []int{}, []int{}}, nil},
		{"* * * 1 * *", Rule{[]int{}, []int{}, []int{}, []int{1}, []int{}, []int{}}, nil},
		{"* * * * 1 *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{1}, []int{}}, nil},
		{"* * * * * 1", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{1}}, nil},

		{"1,2,3 * * * * *", Rule{[]int{1, 2, 3}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* 1,2,3 * * * *", Rule{[]int{}, []int{1, 2, 3}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * 1,2,3 * * *", Rule{[]int{}, []int{}, []int{1, 2, 3}, []int{}, []int{}, []int{}}, nil},
		{"* * * 1,2,3 * *", Rule{[]int{}, []int{}, []int{}, []int{1, 2, 3}, []int{}, []int{}}, nil},
		{"* * * * 1,2,3 *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{1, 2, 3}, []int{}}, nil},
		{"* * * * * 1,2,3", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{1, 2, 3}}, nil},

		{"1-3 * * * * *", Rule{[]int{1, 2, 3}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* 1-3 * * * *", Rule{[]int{}, []int{1, 2, 3}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * 1-3 * * *", Rule{[]int{}, []int{}, []int{1, 2, 3}, []int{}, []int{}, []int{}}, nil},
		{"* * * 1-3 * *", Rule{[]int{}, []int{}, []int{}, []int{1, 2, 3}, []int{}, []int{}}, nil},
		{"* * * * 1-3 *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{1, 2, 3}, []int{}}, nil},
		{"* * * * * 1-3", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{1, 2, 3}}, nil},

		{"*/20 * * * * *", Rule{[]int{0, 20, 40}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* */30 * * * *", Rule{[]int{}, []int{0, 30}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"* * */6 * * *", Rule{[]int{}, []int{}, []int{0, 6, 12, 18}, []int{}, []int{}, []int{}}, nil},
		{"* * * */10 * *", Rule{[]int{}, []int{}, []int{}, []int{10, 20, 30}, []int{}, []int{}}, nil},
		{"* * * * */4 *", Rule{[]int{}, []int{}, []int{}, []int{}, []int{4, 8, 12}, []int{}}, nil},
		{"* * * * * */2", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{0, 2, 4, 6}}, nil},

		{"@weekly", Rule{[]int{0}, []int{0}, []int{0}, []int{}, []int{}, []int{0}}, nil},
		{"@hourly", Rule{[]int{0}, []int{0}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"@daily", Rule{[]int{0}, []int{0}, []int{0}, []int{}, []int{}, []int{}}, nil},
		{"@monthly", Rule{[]int{0}, []int{0}, []int{0}, []int{1}, []int{}, []int{}}, nil},
		{"@yearly", Rule{[]int{0}, []int{0}, []int{0}, []int{1}, []int{1}, []int{}}, nil},

		{"@every second", Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"@every minute", Rule{[]int{0}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"@every hour", Rule{[]int{0}, []int{0}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"@every day", Rule{[]int{0}, []int{0}, []int{0}, []int{}, []int{}, []int{}}, nil},
		{"@every month", Rule{[]int{0}, []int{0}, []int{0}, []int{1}, []int{}, []int{}}, nil},
		{"@every week", Rule{[]int{0}, []int{0}, []int{0}, []int{}, []int{}, []int{0}}, nil},
		{"@every year", Rule{[]int{0}, []int{0}, []int{0}, []int{1}, []int{1}, []int{}}, nil},

		{"@every 20 second", Rule{[]int{0, 20, 40}, []int{}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"@every 30 minute", Rule{[]int{0}, []int{0, 30}, []int{}, []int{}, []int{}, []int{}}, nil},
		{"@every 6 hour", Rule{[]int{0}, []int{0}, []int{0, 6, 12, 18}, []int{}, []int{}, []int{}}, nil},
		{"@every 10 day", Rule{[]int{0}, []int{0}, []int{0}, []int{10, 20, 30}, []int{}, []int{}}, nil},
		{"@every 4 month", Rule{[]int{0}, []int{0}, []int{0}, []int{1}, []int{4, 8, 12}, []int{}}, nil},
		{"@every 2 dayofweek", Rule{[]int{0}, []int{0}, []int{0}, []int{}, []int{}, []int{0, 2, 4, 6}}, nil},
	}

	for _, v := range tests {
		vv, err := Parse(v.in)
		assert.Equal(t, err, v.err, v)
		assert.Equal(t, vv, v.out, v)
	}

	fails := []string{
		"@likexian",
		"0 0",
		"-3 * * * *",
		"3- * * * *",
		"x-3 * * * *",
		"3-x * * * *",
		"100-1000 * * * *",
		"1000-100 * * * *",
		"1/x * * * *",
		"100/x * * * *",
		"x/100 * * * *",
		"x * * * *",
		"1000 * * * *",
		"1,1000 * * * *",
		"1,x,3 * * * *",
		"@every x",
		"@every x second",
		"* * * * janx *",
		"* * * * janx-mar *",
		"* * * * janx,feb,mar *",
		"* * * * * sunx",
		"* * * * * sunx-tue",
		"* * * * * sunx,mon,tue",
	}

	for _, v := range fails {
		_, err := Parse(v)
		assert.NotNil(t, err, v)
	}
}

func TestMustParse(t *testing.T) {
	assert.Panic(t, func() { MustParse("@every night") })
	assert.NotPanic(t, func() { MustParse("@every second") })
}

func TestIsDue(t *testing.T) {
	tests := []struct {
		now  time.Time
		rule Rule
		out  bool
	}{
		{time.Date(2019, 04, 10, 0, 0, 0, 0, time.UTC), Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}}, true},
		{time.Date(2019, 04, 10, 0, 0, 0, 0, time.UTC), Rule{[]int{0}, []int{}, []int{}, []int{}, []int{}, []int{}}, true},
		{time.Date(2019, 04, 10, 0, 0, 0, 0, time.UTC), Rule{[]int{0, 1}, []int{}, []int{}, []int{}, []int{}, []int{}}, true},
		{time.Date(2019, 04, 10, 0, 0, 0, 0, time.UTC), Rule{[]int{1}, []int{}, []int{}, []int{}, []int{}, []int{}}, false},
		{time.Date(2019, 04, 10, 0, 0, 0, 0, time.UTC), Rule{[]int{1, 2}, []int{}, []int{}, []int{}, []int{}, []int{}}, false},
	}

	for _, v := range tests {
		assert.Equal(t, isDue(v.now, v.rule), v.out)
	}
}
