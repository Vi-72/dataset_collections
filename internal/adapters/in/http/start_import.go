package http

import (
	"context"
	"dataset-collections/internal/generated/servers"
)

func (h *ApiHandler) StartImport(ctx context.Context, _ servers.StartImportRequestObject) (servers.StartImportResponseObject, error) {
	result, err := h.startImportHandler.Handle(ctx)
	if err != nil {
		return servers.StartImport500Response{}, nil
	}

	return servers.StartImport202JSONResponse{
		JobId:  result.JobID,
		Status: "pending",
	}, nil
}
