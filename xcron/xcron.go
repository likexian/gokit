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
	"fmt"
	"strconv"
	"strings"
)

// Field type of rule
const (
	Second = iota
	Minute
	Hour
	DayOfMonth
	Month
	DayOfWeek
)

// Rule is parsed cron rule
type Rule struct {
	Second     []int
	Minute     []int
	Hour       []int
	DayOfMonth []int
	Month      []int
	DayOfWeek  []int
}

// Version returns package version
func Version() string {
	return "0.2.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Parse parse single cron rule
// Base on https://en.wikipedia.org/wiki/Cron
// Fields: second minute hour dayOfMonth month dayOfWeek
//         *      *      *    *          *     *
func Parse(s string) (r Rule, err error) {
	r = Rule{[]int{}, []int{}, []int{}, []int{}, []int{}, []int{}}

	if s == "" || s == "*" {
		return
	}

	s = strings.TrimSpace(s)
	if s[0] == '@' {
		s, err = parseMacros(s)
		if err != nil {
			return
		}
	}

	fs := strings.Fields(s)
	if len(fs) < 6 {
		fs = append([]string{"0"}, fs...)
	}

	if len(fs) != 6 {
		return r, fmt.Errorf("xcron: unrecognized rule: %s", s)
	}

	for i := 0; i < len(fs); i++ {
		err := r.parseField(fs[i], i)
		if err != nil {
			return r, err
		}
	}

	return
}

// parseField parse every fields
func (r *Rule) parseField(s string, t int) (err error) {
	switch t {
	case Second:
		if strings.Contains(s, ",") {
			r.Second, err = getField(s, 0, 59)
		} else {
			r.Second, err = getRange(s, 0, 59)
		}
	case Minute:
		if strings.Contains(s, ",") {
			r.Minute, err = getField(s, 0, 59)
		} else {
			r.Minute, err = getRange(s, 0, 59)
		}
	case Hour:
		if strings.Contains(s, ",") {
			r.Hour, err = getField(s, 0, 23)
		} else {
			r.Hour, err = getRange(s, 0, 23)
		}
	case DayOfMonth:
		if strings.Contains(s, ",") {
			r.DayOfMonth, err = getField(s, 1, 31)
		} else {
			r.DayOfMonth, err = getRange(s, 1, 31)
		}
	case Month:
		if strings.Contains(s, ",") {
			r.Month, err = getField(s, 1, 12)
		} else {
			r.Month, err = getRange(s, 1, 12)
		}
	case DayOfWeek:
		if strings.Contains(s, ",") {
			r.DayOfWeek, err = getField(s, 0, 6)
		} else {
			r.DayOfWeek, err = getRange(s, 0, 6)
		}
	}

	return err
}

// getRange get int array from string range, for example 3, 0-23, */3
func getRange(s string, min, max int) ([]int, error) {
	r := []int{}

	if s == "*" {
		return r, nil
	}

	if strings.Contains(s, "-") {
		ss := strings.Split(s, "-")
		sl, err := strconv.Atoi(ss[0])
		if err != nil {
			return r, fmt.Errorf("xcron: unrecognized charset: %s", ss[0])
		}
		sr, err := strconv.Atoi(ss[1])
		if err != nil {
			return r, fmt.Errorf("xcron: unrecognized charset: %s", ss[1])
		}
		if sl > sr {
			st := sr
			sr = sl
			sl = st
		}
		if sl < min || sr > max {
			return r, fmt.Errorf("xcron: %d is not in [%d, %d]", sr, min, max)
		}
		for i := sl; i <= sr; i++ {
			r = append(r, i)
		}
	} else if strings.Contains(s, "/") {
		ss := strings.Split(s, "/")
		sr, err := strconv.Atoi(ss[1])
		if err != nil {
			return r, fmt.Errorf("xcron: unrecognized charset: %s", ss[1])
		}
		if sr < min || sr > max {
			return r, fmt.Errorf("xcron: %d is not in [%d, %d]", sr, min, max)
		}
		for i := min; i <= max; i++ {
			if i%sr == 0 {
				r = append(r, i)
			}
		}
	} else {
		sr, err := strconv.Atoi(s)
		if err != nil {
			return r, fmt.Errorf("xcron: unrecognized charset: %s", s)
		}
		if sr < min || sr > max {
			return r, fmt.Errorf("xcron: %d is not in [%d, %d]", sr, min, max)
		}
		r = append(r, sr)
	}

	return r, nil
}

// getField get int array from string fields, for example 0,1,2
func getField(s string, min, max int) ([]int, error) {
	r := []int{}

	for _, v := range strings.Split(s, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			vv, err := strconv.Atoi(v)
			if err != nil {
				return r, fmt.Errorf("xcron: unrecognized charset: %s", v)
			}
			if vv < min || vv > max {
				return r, fmt.Errorf("xcron: %d is not in [%d, %d]", vv, min, max)
			}
			r = append(r, vv)
		}
	}

	return r, nil
}

// parseMacros parse nonstandard predefined scheduling definitions
// returns as standard scheduling definitions
func parseMacros(s string) (string, error) {
	switch strings.ToLower(s) {
	case "@yearly", "@annually":
		return "0 0 1 1 *", nil
	case "@monthly":
		return "0 0 1 * *", nil
	case "@daily", "@midnight":
		return "0 0 * * *", nil
	case "@hourly":
		return "0 * * * *", nil
	case "@weekly":
		return "0 0 * * 0", nil
	default:
		every := "@every "
		if strings.HasPrefix(s, every) {
			s = strings.TrimSpace(s[len(every):])
			ss := strings.Fields(s)
			// @every hour
			ev := "*"
			vv := 0
			if len(ss) > 1 {
				vv, err := strconv.Atoi(ss[0])
				if err == nil && vv > 0 {
					// @every 2 hour -> */2
					ev = "*/" + ss[0]
					ss[0] = ss[1]
				} else {
					return "", fmt.Errorf("xcron: unrecognized macros: %s", s)
				}
			}
			switch ss[0] {
			case "year":
				if ev == "*" {
					return "0 0 1 1 *", nil
				}
			case "month":
				if vv <= 12 {
					return fmt.Sprintf("0 0 1 %s *", ev), nil
				}
			case "day":
				if vv <= 31 {
					return fmt.Sprintf("0 0 %s * *", ev), nil
				}
			case "hour":
				if vv <= 24 {
					return fmt.Sprintf("0 %s * * *", ev), nil
				}
			case "minute":
				if vv <= 60 {
					return fmt.Sprintf("%s * * * *", ev), nil
				}
			case "second":
				if vv <= 60 {
					return fmt.Sprintf("%s * * * * *", ev), nil
				}
			case "week":
				if ev == "*" {
					return "0 0 0 * * 0", nil
				}
			case "dayofweek":
				if vv < 7 {
					return fmt.Sprintf("0 0 0 * * %s", ev), nil
				}
			}
		}
		return "", fmt.Errorf("xcron: unrecognized macros: %s", s)
	}
}
