package commands

import (
	"context"
	"task-processing-service/internal/core/domain/model/task"
	"task-processing-service/internal/core/ports"
	"task-processing-service/internal/pkg/errs"
)

type ExecuteTaskCommandHandler interface {
	Handle(ctx context.Context, t *task.Task) error
}

type executeTaskCommandHandler struct {
	execService ports.TaskExecutionService
	repo        ports.TaskRepository
}

func NewExecuteTaskCommandHandler(
	execService ports.TaskExecutionService,
	repo ports.TaskRepository,
) (ExecuteTaskCommandHandler, error) {
	if execService == nil {
		return nil, errs.NewValueIsRequiredError("execution service")
	}
	if repo == nil {
		return nil, errs.NewValueIsRequiredError("task repository")
	}
	return &executeTaskCommandHandler{
		execService: execService,
		repo:        repo,
	}, nil
}

func (h *executeTaskCommandHandler) Handle(ctx context.Context, t *task.Task) error {
	if t == nil {
		return errs.NewValueIsRequiredError("task")
	}

	if err := h.repo.Save(ctx, t); err != nil {
		return err
	}

	return h.execService.Execute(ctx, t)
}
