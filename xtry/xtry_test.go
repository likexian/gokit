/*
 * Copyright 2012-2026 Li Kexian
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
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestSucc(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{}
	err := c.Run(ctx, func(context.Context) error { return nil })
	assert.Nil(t, err)
}

func TestCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(500*time.Millisecond, cancel)

	c := Config{}
	err := c.Run(ctx, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, Cancelled)
	assert.Equal(t, e.Times, 1)
}

func TestTimeout(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		Timeout: 1 * time.Second,
	}
	err := c.Run(ctx, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, Timeout)
	assert.Equal(t, e.Times, 1)
}

func TestMaxTries(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		MaxTries: 1,
	}
	err := c.Run(ctx, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, MaxTries)
	assert.Equal(t, e.Times, 1)
}

func TestRetryDelay(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		RetryDelay: func() time.Duration { return 500 * time.Millisecond },
		Timeout:    1 * time.Second,
	}
	err := c.Run(ctx, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, Timeout)
	assert.Equal(t, e.Times, 2)
}

func TestShouldRetry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		ShouldRetry: func(error) bool { return true },
		Timeout:     1 * time.Second,
	}
	err := c.Run(ctx, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, Timeout)
	assert.Equal(t, e.Times, 1)
}

func TestNonShouldRetry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		ShouldRetry: func(error) bool { return false },
	}
	err := c.Run(ctx, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, NonRetry)
	assert.Equal(t, e.Times, 0)
}

func TestRetryableError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		Timeout: 1 * time.Second,
	}
	err := c.Run(ctx, func(context.Context) error { return RetryableError(nil) })
	assert.Nil(t, err)

	err = c.Run(ctx, func(context.Context) error { return RetryableError(fmt.Errorf("RetryableError")) })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, Timeout)
	assert.Equal(t, e.Times, 1)
}

func TestNonRetryableError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := Config{
		Timeout: 1 * time.Second,
	}
	err := c.Run(ctx, func(context.Context) error { return NonRetryableError(nil) })
	assert.Nil(t, err)

	err = c.Run(ctx, func(context.Context) error { return NonRetryableError(fmt.Errorf("NonRetryableError")) })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, NonRetry)
	assert.Equal(t, e.Times, 0)
}

func TestRetry(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	err := Retry(ctx, 1*time.Second, func(context.Context) error { return fmt.Errorf("error") })
	assert.NotNil(t, err)

	var e *RetryExhaustedError
	ok := errors.As(err, &e)
	assert.True(t, ok)
	assert.Equal(t, e.Type, Timeout)
	assert.Equal(t, e.Times, 1)
}

func TestRetryExhaustedError(t *testing.T) {
	t.Parallel()

	err := RetryExhaustedError{Err: nil, Type: Timeout}
	assert.Equal(t, err.Error(), "<nil>")

	err = RetryExhaustedError{Err: fmt.Errorf("error"), Type: Timeout}
	assert.Equal(t, err.Error(), "xtry: retry exhausted, type: Timeout, error: error")
}

func TestRetryError(t *testing.T) {
	t.Parallel()

	err := &RetryError{Err: nil, Retryable: true}
	assert.Equal(t, err.Error(), "<nil>")

	err = &RetryError{Err: fmt.Errorf("error"), Retryable: true}
	assert.Equal(t, err.Error(), "xtry: retryable error: error")

	err = &RetryError{Err: fmt.Errorf("error"), Retryable: false}
	assert.Equal(t, err.Error(), "xtry: nonretryable error: error")
}
