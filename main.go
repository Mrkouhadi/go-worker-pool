package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	numTasks    = 100
	workerCount = 8
	rateLimit   = 100 // ms
	taskTimeout = 5   // seconds
	retryCount  = 3
)

func main() {

	pool := NewWorkerPool(workerCount, time.Millisecond*rateLimit, numTasks)
	pool.Start()

	results := make(chan Result, numTasks)
	tasks := make([]Task, 0, numTasks)

	for i := 0; i < numTasks; i++ {
		id := i
		localID := id
		tasks = append(tasks, Task{
			ID:       localID,
			Priority: rand.Intn(5),
			Timeout:  time.Second * taskTimeout,
			Retries:  retryCount,
			ResultCh: results,
			Exec: func(ctx context.Context) (any, error) {
				if rand.Float32() < 0.3 {
					return nil, errors.New("random error")
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
				return fmt.Sprintf("Task-%d done", localID), nil
			},
		})
	}

	pool.Submit(tasks)

	go func() {
		for r := range results {
			if r.Err != nil {
				fmt.Printf("Task %d failed: %v\n", r.TaskID, r.Err)
			} else {
				fmt.Printf("Task %d result: %v\n", r.TaskID, r.Value)
			}
		}
	}()

	pool.Stop(results)

	fmt.Printf("\nTotal dropped tasks: %d\n", pool.droppedTaskCount)
}
