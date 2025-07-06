package task

import (
	"context"
	"errors"
	"task-processing-service/internal/core/domain/model/task"
	taskports "task-processing-service/internal/core/ports"
	"task-processing-service/internal/pkg/errs"
)

type executionService struct {
	runner *Runner
}

var (
	ErrTaskAlreadyExecuted = errors.New("task is already completed or failed")
)

func NewExecutionService(runner *Runner) taskports.TaskExecutionService {
	return &executionService{runner: runner}
}

func (s *executionService) Execute(ctx context.Context, t *task.Task) error {
	if t == nil {
		return errs.NewValueIsRequiredError("task")
	}

	if t.Status() == task.StatusCompleted || t.Status() == task.StatusFailed {
		return ErrTaskAlreadyExecuted
	}

	if err := s.runner.Enqueue(t); err != nil {
		return err
	}

	return nil
}
