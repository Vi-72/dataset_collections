package queries

import (
	"context"
	"github.com/google/uuid"
	"task-processing-service/internal/core/ports"
	"task-processing-service/internal/pkg/errs"
)

type DeleteTaskQueryHandler interface {
	Handle(ctx context.Context, id uuid.UUID) error
}

type deleteTaskQueryHandler struct {
	repo ports.TaskRepository
}

func NewDeleteTaskQueryHandler(repo ports.TaskRepository) DeleteTaskQueryHandler {
	return &deleteTaskQueryHandler{repo: repo}
}

func (h *deleteTaskQueryHandler) Handle(ctx context.Context, id uuid.UUID) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return errs.NewObjectNotFoundErrorWithCause("task", id.String(), err)
	}
	return nil
}
