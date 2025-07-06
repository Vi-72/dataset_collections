package task

import (
	"fmt"
	"math/rand"
	"task-processing-service/internal/core/domain/model/task"
	"task-processing-service/internal/pkg/errs"
	"time"
)

type Runner struct {
	taskQueue chan *task.Task
	workers   int
	quit      chan struct{}
}

func NewTaskRunner(workers int, queueSize int) *Runner {
	runner := &Runner{
		taskQueue: make(chan *task.Task, queueSize),
		workers:   workers,
		quit:      make(chan struct{}),
	}
	runner.start()
	return runner
}

func (r *Runner) start() {
	for i := 0; i < r.workers; i++ {
		go func(workerID int) {
			for {
				select {
				case t := <-r.taskQueue:
					r.executeTask(t, workerID)
				case <-r.quit:
					return
				}
			}
		}(i)
	}
}

func (r *Runner) executeTask(t *task.Task, workerID int) {
	if t == nil {
		return
	}
	if t.Status() == task.StatusCompleted || t.Status() == task.StatusFailed {
		return
	}

	t.MarkRunning()

	delay := time.Duration(3+rand.Intn(3)) * time.Minute
	time.Sleep(delay)

	if rand.Float64() < 0.8 {
		t.Complete(fmt.Sprintf("executed by worker %d", workerID), delay)
	} else {
		t.Fail("execution failed", delay)
	}
}

func (r *Runner) Enqueue(t *task.Task) error {
	if t == nil {
		return errs.NewValueIsRequiredError("task")
	}
	select {
	case r.taskQueue <- t:
		return nil
	default:
		return errs.NewOperationFailedError("task queue is full")
	}
}

func (r *Runner) Stop() {
	close(r.quit)
}
