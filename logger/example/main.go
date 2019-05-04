/*
 * Go module for doing logging
 * https://www.likexian.com/
 *
 * Copyright 2015-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package main

import (
	"github.com/likexian/logger-go"
	"time"
)

func main() {
	sizeLog()
	dailyLog()
}

// dailyLog do log rotate daily
func dailyLog() {
	log, err := logger.File("test.daily.log", logger.DEBUG)
	if err != nil {
		panic(err)
	}

	log.SetDailyRotate(3)
	for {
		log.Info("This is a test log")
		time.Sleep(1 * time.Second)
	}
}

// sizeLog do log rotate by size
func sizeLog() {
	log, err := logger.File("test.size.log", logger.DEBUG)
	if err != nil {
		panic(err)
	}

	log.SetSizeRotate(10, 1000000)
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				log.Info("This is a test log")
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
