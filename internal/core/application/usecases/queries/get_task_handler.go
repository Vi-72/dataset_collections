package queries

import (
	"context"
	"github.com/google/uuid"
	"task-processing-service/internal/core/domain/model/task"
	"task-processing-service/internal/core/ports"
	"task-processing-service/internal/pkg/errs"
)

type GetTaskQueryHandler interface {
	Handle(ctx context.Context, id uuid.UUID) (*task.Task, error)
}

type getTaskQueryHandler struct {
	repo ports.TaskRepository
}

func NewGetTaskQueryHandler(repo ports.TaskRepository) GetTaskQueryHandler {
	return &getTaskQueryHandler{repo: repo}
}

func (h *getTaskQueryHandler) Handle(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	t, err := h.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errs.NewObjectNotFoundErrorWithCause("task", id.String(), err)
	}
	return t, nil
}
