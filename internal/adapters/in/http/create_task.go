package http

import (
	"context"
	"task-processing-service/internal/core/domain/model/task"
	"task-processing-service/internal/generated/servers"
	"task-processing-service/internal/pkg/errs"
	"time"

	"github.com/google/uuid"
)

func (h *TaskHandler) CreateTask(
	ctx context.Context,
	_ servers.CreateTaskRequestObject,
) (servers.CreateTaskResponseObject, error) {
	id := uuid.New()
	t, err := task.NewTask(id, time.Now())
	if err != nil {
		return nil, errs.NewInternalServerError("failed to create task: %w", err)
	}

	err = h.executeHandler.Handle(ctx, t)
	if err != nil {
		return nil, errs.NewInternalServerError("failed to execute task: %w", err)
	}

	resp := servers.CreateTask201JSONResponse{
		Id:        t.ID().String(),
		Status:    servers.Status(t.Status()),
		CreatedAt: t.CreatedAt(),
	}

	return resp, nil
}
