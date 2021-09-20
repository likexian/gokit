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

package xtry

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	// Timeout is retry exhausted timeout
	Timeout ExhaustedType = "Timeout"
	// MaxTries is retry exhausted max times
	MaxTries ExhaustedType = "MaxTries"
	// Cancelled is retry is cancelled
	Cancelled ExhaustedType = "Cancelled"
	// NonRetry is non retryable
	NonRetry ExhaustedType = "NonRetry"
)

// ExhaustedType is retry exhausted type
type ExhaustedType string

// Config represents a retry config
type Config struct {
	// Retry until timeout elapsed, 0 means forever
	Timeout time.Duration

	// MaxTries is max retry times, 0 means forever
	MaxTries int

	// RetryDelay returns dealy time after failed, default is 1s
	RetryDelay func() time.Duration

	// ShouldRetry returns wether error should be retried, default true
	ShouldRetry func(error) bool
}

// Version returns package version
func Version() string {
	return "0.4.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Run calls fn util ctx is cancelled or max retry exhausted
func (c Config) Run(ctx context.Context, fn func(context.Context) error) error {
	retryDelay := func() time.Duration { return 1 * time.Second }
	if c.RetryDelay != nil {
		retryDelay = c.RetryDelay
	}

	shouldRetry := func(error) bool { return true }
	if c.ShouldRetry != nil {
		shouldRetry = c.ShouldRetry
	}

	var timeout <-chan time.Time
	if c.Timeout != 0 {
		timeout = time.After(c.Timeout)
	}

	var err error
	for try := 0; ; try++ {
		if c.MaxTries != 0 && try == c.MaxTries {
			return &RetryExhaustedError{Err: err, Type: MaxTries, Times: try}
		}
		if err = fn(ctx); err == nil {
			return nil
		}
		var e *RetryError
		if ok := errors.As(err, &e); ok {
			if e == nil {
				return nil
			}
			if !e.Retryable {
				return &RetryExhaustedError{Err: e.Err, Type: NonRetry, Times: try}
			}
		} else {
			if !shouldRetry(err) {
				return &RetryExhaustedError{Err: err, Type: NonRetry, Times: try}
			}
		}
		select {
		case <-ctx.Done():
			return &RetryExhaustedError{Err: err, Type: Cancelled, Times: try}
		case <-timeout:
			return &RetryExhaustedError{Err: err, Type: Timeout, Times: try}
		default:
			time.Sleep(retryDelay())
		}
	}
}

// Retry do retry with timeout
func Retry(ctx context.Context, timeout time.Duration, fn func(context.Context) error) error {
	return Config{Timeout: timeout}.Run(ctx, fn)
}

// RetryExhaustedError is max retry exhausted error
type RetryExhaustedError struct {
	Err   error
	Type  ExhaustedType
	Times int
}

// Error returns string of max retry exhausted error
func (err *RetryExhaustedError) Error() string {
	if err == nil || err.Err == nil {
		return "<nil>"
	}

	return fmt.Sprintf("xtry: retry exhausted, type: %s, error: %s", err.Type, err.Err)
}

// RetryError is an error with retryable info
type RetryError struct {
	Err       error
	Retryable bool
}

// Error returns string of retry error
func (err *RetryError) Error() string {
	if err == nil || err.Err == nil {
		return "<nil>"
	}

	if err.Retryable {
		return fmt.Sprintf("xtry: retryable error: %s", err.Err)
	}

	return fmt.Sprintf("xtry: nonretryable error: %s", err.Err)
}

// RetryableError returns a retryable error
func RetryableError(err error) *RetryError {
	if err == nil {
		return nil
	}

	return &RetryError{Err: err, Retryable: true}
}

// NonRetryableError returns a not retryable error
func NonRetryableError(err error) *RetryError {
	if err == nil {
		return nil
	}

	return &RetryError{Err: err, Retryable: false}
}
