package http

import (
	"context"
	"github.com/google/uuid"
	"task-processing-service/internal/generated/servers"
	"task-processing-service/internal/pkg/errs"
)

func (h *TaskHandler) DeleteTask(
	ctx context.Context,
	request servers.DeleteTaskRequestObject,
) (servers.DeleteTaskResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, errs.ErrValueIsInvalid
	}

	err = h.deleteHandler.Handle(ctx, id)
	if err != nil {
		if errs.IsNotFound(err) {
			return servers.DeleteTask404Response{}, nil
		}
		return nil, errs.NewInternalServerError("failed to delete task: %v", err)
	}

	return servers.DeleteTask204Response{}, nil
}
