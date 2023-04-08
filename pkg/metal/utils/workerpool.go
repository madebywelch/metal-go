package utils

import "sync"

// Task is a function type that represents a task to be performed by a worker.
type Task func()

// WorkerPool is a pool of workers that can process tasks concurrently.
type WorkerPool struct {
	workerCount int
	taskQueue   chan Task
	wg          sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool with a specified number of workers.
func NewWorkerPool(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		workerCount: workerCount,
		taskQueue:   make(chan Task),
	}

	pool.startWorkers()

	return pool
}

// startWorkers initializes and starts the workers.
func (wp *WorkerPool) startWorkers() {
	wp.wg.Add(wp.workerCount)

	for i := 0; i < wp.workerCount; i++ {
		go func() {
			defer wp.wg.Done()

			for task := range wp.taskQueue {
				task()
			}
		}()
	}
}

// AddTask adds a task to the worker pool.
func (wp *WorkerPool) AddTask(task Task) {
	wp.taskQueue <- task
}

// CloseAndWait waits for all tasks to complete and then closes the task queue.
func (wp *WorkerPool) CloseAndWait() {
	close(wp.taskQueue)
	wp.wg.Wait()
}
