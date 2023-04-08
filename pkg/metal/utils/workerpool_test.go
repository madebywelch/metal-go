package utils

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	workerCount := 5
	taskCount := 20

	var completedTasks int32
	wp := NewWorkerPool(workerCount)

	for i := 0; i < taskCount; i++ {
		wp.AddTask(func() {
			time.Sleep(100 * time.Millisecond)
			atomic.AddInt32(&completedTasks, 1)
		})
	}

	wp.CloseAndWait()

	if int(completedTasks) != taskCount {
		t.Errorf("Expected %d completed tasks, but got %d", taskCount, completedTasks)
	}
}

func TestWorkerPoolNoTasks(t *testing.T) {
	workerCount := 5
	wp := NewWorkerPool(workerCount)

	// No tasks added
	wp.CloseAndWait()
}
