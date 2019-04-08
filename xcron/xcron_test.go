/*
 * Copyright 2012-2019 Li Kexian
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
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
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
		"1,x,3 * * * *",
		"@every x",
		"@every x second",
	}

	for _, v := range fails {
		_, err := Parse(v)
		assert.NotNil(t, err, v)
	}
}
