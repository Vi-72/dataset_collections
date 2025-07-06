package http

import (
	"context"
	"github.com/google/uuid"
	"task-processing-service/internal/generated/servers"
	"task-processing-service/internal/pkg/errs"
)

func (h *TaskHandler) GetTask(
	ctx context.Context,
	request servers.GetTaskRequestObject,
) (servers.GetTaskResponseObject, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, errs.ErrValueIsInvalid
	}

	t, err := h.getTaskQuery.Handle(ctx, id)
	if err != nil {
		if errs.IsNotFound(err) {
			return servers.GetTask404Response{}, nil
		}
		return nil, errs.NewInternalServerError("failed to get task: %v", err)
	}
	var durationStr *string
	if t.Duration() != nil {
		d := t.Duration().String() // time.Duration → string
		durationStr = &d
	}

	resp := servers.GetTask200JSONResponse{
		Id:        t.ID().String(),
		Status:    servers.Status(t.Status()),
		CreatedAt: t.CreatedAt(),
		Result:    t.Result(),
		Error:     t.ErrorMessage(),
		Duration:  durationStr,
	}
	return resp, nil
}
