package main

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Task defines a unit of work to be processed by the worker pool.
type Task struct {
	ID       int
	Priority int
	Exec     func(ctx context.Context) (any, error)
	Timeout  time.Duration
	Retries  int
	ResultCh chan Result
}

// Result represents the outcome of a task.
type Result struct {
	TaskID int
	Value  any
	Err    error
}

// WorkerPool manages a pool of goroutines to execute submitted tasks.
type WorkerPool struct {
	workerCount      int
	rateLimit        time.Duration
	tasks            chan Task
	wg               sync.WaitGroup
	droppedTaskCount int
	resultsWg        sync.WaitGroup
	mutex            sync.Mutex
}

func NewWorkerPool(workerCount int, rateLimit time.Duration, queueSize int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		rateLimit:   rateLimit,
		tasks:       make(chan Task, queueSize),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) Stop(results chan Result) {
	close(wp.tasks)
	wp.wg.Wait()
	wp.resultsWg.Wait()
	close(results) // owned and closed only after all writers finish
}

func (wp *WorkerPool) Submit(tasks []Task) {
	// sort tasks based on the priority
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority
	})

	for _, task := range tasks {
		select {
		case wp.tasks <- task:
		default:
			wp.mutex.Lock()
			wp.droppedTaskCount++
			wp.mutex.Unlock()
			fmt.Printf("Task %d dropped due to queue overflow\n", task.ID)
		}
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.tasks {
		time.Sleep(wp.rateLimit)
		wp.resultsWg.Add(1)
		go wp.handleTask(task, id)
	}
}

func (wp *WorkerPool) handleTask(task Task, workerID int) {
	defer wp.resultsWg.Done()

	var result Result
	for attempt := 0; attempt <= task.Retries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), task.Timeout)
		val, err := task.Exec(ctx)
		cancel()

		if err == nil {
			result = Result{TaskID: task.ID, Value: val, Err: nil}
			fmt.Printf("Worker %d: Task %d success\n", workerID, task.ID)
			break
		} else {
			result = Result{TaskID: task.ID, Value: nil, Err: err}
			// creates exponentially increasing delays between retry attempts.
			sleep := time.Duration(100*(1<<attempt)) * time.Millisecond // 1st try wait 100ms, 2nd retry wait 200ms
			// Adds a random jitter between 0-50ms to the sleep duration. it Helps prevent the "thundering herd" problem when many workers retry simultaneously
			sleep += time.Duration(rand.Intn(50)) * time.Millisecond
			time.Sleep(sleep)
		}
	}

	if task.ResultCh != nil {
		task.ResultCh <- result
	}
}
