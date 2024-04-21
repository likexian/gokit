/*
 * Copyright 2012-2024 Li Kexian
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

package main

import (
	"time"

	"github.com/likexian/gokit/xlog"
)

func main() {
	sizeLog()
	dailyLog()
}

// dailyLog do log rotate daily
func dailyLog() {
	log, err := xlog.File("test.daily.log", xlog.DEBUG)
	if err != nil {
		panic(err)
	}

	_ = log.SetDailyRotate(3)
	for {
		log.Info("This is a test log")
		time.Sleep(1 * time.Second)
	}
}

// sizeLog do log rotate by size
func sizeLog() {
	log, err := xlog.File("test.size.log", xlog.DEBUG)
	if err != nil {
		panic(err)
	}

	_ = log.SetSizeRotate(10, 1000000)
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				log.Info("This is a test log")
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
