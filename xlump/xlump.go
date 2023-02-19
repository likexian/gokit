/*
 * Copyright 2012-2023 Li Kexian
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

package xlump

import (
	"runtime"
	"sync"
)

//             .---------.                   .---------.                    .---------.
// Task.Add -> |         |                   |         |                    |         |
// Task.Add -> |         |    Worker.Work    |         |                    |         |
// Task.Add -> | TaskIn  | -> Worker.Work -> | TaskOut | -> Merger.Merge -> | TaskSum |
// Task.Add -> |         |    Worker.Work    |         |                    |         |
// Task.Add -> |         |                   |         |                    |         |
//             '---------'                   '---------'                    '---------'

// Task is task put to queue
type Task interface{}

// Work work with the task
type Work func(Task) Task

// Merge merge worker result
type Merge func(Task, Task) Task

// Queue is the work queue
type Queue struct {
	TaskIn   chan Task
	TaskOut  chan Task
	TaskSum  chan Task
	WorkWait sync.WaitGroup
}

// Version returns package version
func Version() string {
	return "0.3.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// New returns new work queue
func New(bufferSize int) *Queue {
	if bufferSize <= 0 {
		bufferSize = 0
	}

	return &Queue{
		TaskIn:   make(chan Task, bufferSize),
		TaskOut:  make(chan Task, bufferSize),
		TaskSum:  make(chan Task),
		WorkWait: sync.WaitGroup{},
	}
}

// SetWorker start worker to do work
func (q *Queue) SetWorker(work Work, number int) *Queue {
	if number <= 0 {
		number = runtime.NumCPU()
	}

	for i := 0; i < number; i++ {
		go q.worker(work)
		q.WorkWait.Add(1)
	}

	return q
}

// worker get task and work
func (q *Queue) worker(work Work) {
	for {
		t, ok := <-q.TaskIn
		if !ok {
			q.WorkWait.Done()
			return
		}
		q.TaskOut <- work(t)
	}
}

// SetMerger start merger to do merge
func (q *Queue) SetMerger(merge Merge, result Task) *Queue {
	go func() {
		for {
			t, ok := <-q.TaskOut
			if !ok {
				q.TaskSum <- result
				return
			}
			result = merge(result, t)
		}
	}()

	return q
}

// Add add task to queue
func (q *Queue) Add(task Task) {
	q.TaskIn <- task
}

// Wait close in queue and wait for result
func (q *Queue) Wait() Task {
	close(q.TaskIn)

	q.WorkWait.Wait()
	close(q.TaskOut)

	r := <-q.TaskSum
	close(q.TaskSum)

	return r
}
