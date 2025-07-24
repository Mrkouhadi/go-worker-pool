package main

import (
	"fmt"
	"time"
)

const (
	workerCount = 8
	rateLimit   = time.Millisecond * 100
)

func main() {

	pool := NewWorkerPool(workerCount, rateLimit, len(tasks))
	pool.Start()

	results := make(chan Result, len(tasks))

	// Set ResultCh for each task before submission
	for i := range tasks {
		tasks[i].ResultCh = results
	}
	// submit tasks
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
