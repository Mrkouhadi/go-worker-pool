package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Task struct {
	ID      int
	Retries int
	Timeout time.Duration
	Exec    func(cxt context.Context) error
}
type Result struct {
	TaskID int
	Value  interface{}
	Err    error
}
type WorkerPool struct {
	Tasks            chan Task
	WorkersCount     int
	Results          chan Result
	RateLimit        time.Duration
	droppedTaskCount int64
}

func LaunchWorkerPool(WorkersCount int, RateLimit time.Duration) *WorkerPool {
	wp := &WorkerPool{
		Tasks:            make(chan Task, 150),
		WorkersCount:     WorkersCount,
		Results:          make(chan Result, 150),
		RateLimit:        RateLimit,
		droppedTaskCount: 0,
	}
	wp.StartWorkers()
	return wp
}
func (wp *WorkerPool) Submit(task Task, timeout time.Duration) {
	select {
	case wp.Tasks <- task:
	case <-time.After(timeout):
		atomic.AddInt64(&wp.droppedTaskCount, 1)
		wp.Results <- Result{
			TaskID: task.ID,
			Err:    errors.New("task submit timeout"),
		}
	}
}

// Worker loop that applies timeout + retries to each task
func (wp *WorkerPool) StartWorkers() {
	for i := 0; i < wp.WorkersCount; i++ {
		go func(workerID int) {
			for task := range wp.Tasks {
				var lastErr error
				for attempt := 0; attempt <= task.Retries; attempt++ {
					ctx, cancel := context.WithTimeout(context.Background(), task.Timeout)
					err := task.Exec(ctx)
					cancel()

					if err == nil {
						wp.Results <- Result{TaskID: task.ID, Value: "✅ success", Err: nil}
						break
					} else {
						// creates exponentially increasing delays between retry attempts.
						sleep := time.Duration(100*(1<<attempt)) * time.Millisecond // 1st try wait 100ms, 2nd retry wait 200ms
						// Adds a random jitter between 0-50ms to the sleep duration. it Helps prevent the "thundering herd" problem when many workers retry simultaneously
						sleep += time.Duration(rand.Intn(50)) * time.Millisecond
						time.Sleep(sleep)
						lastErr = err
					}

				}
				if lastErr != nil {
					wp.Results <- Result{TaskID: task.ID, Err: lastErr}
				}
				// apply rate limit
				time.Sleep(wp.RateLimit)
			}
		}(i)
	}
}

var taskSubmissionTimeout time.Duration = time.Second * 3

func main() {
	rand.Seed(time.Now().UnixNano()) // to get non-deterministic jitter on the first run
	wp := LaunchWorkerPool(10, time.Millisecond*200)
	// Result listener
	go func() {
		for res := range wp.Results {
			if res.Err != nil {
				fmt.Printf("❌ Task #%d failed: %v\n", res.TaskID, res.Err)
			} else {
				fmt.Printf("✅ Task #%d completed: %v\n", res.TaskID, res.Value)
			}
		}

	}()
	// simulate submission of tasks by my http handlers
	go func() {
		i := 1
		for {
			id := i
			task := Task{
				ID:      id,
				Retries: 2,
				Timeout: time.Second * 2,
				Exec: func(ctx context.Context) error {
					fmt.Printf("🚀 Processing Task #%d\n", id)
					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(time.Duration(500+randInt(0, 1500)) * time.Millisecond):
						// simulate random delay
						if randInt(0, 10) < 2 {
							return errors.New("simulated task error")
						}
						return nil
					}
				},
			}
			wp.Submit(task, taskSubmissionTimeout)
			i++
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 3)
			dropped := atomic.LoadInt64(&wp.droppedTaskCount)
			fmt.Printf("%d tasks have been dropped!❌❌❌❌❌❌❌❌❌❌❌\n", dropped)
		}
	}()
	select {}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
