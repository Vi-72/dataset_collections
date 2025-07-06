package http

import (
	"context"
	"dataset-collections/internal/generated/servers"
)

func (h *ApiHandler) StartImport(ctx context.Context, _ servers.StartImportRequestObject) (servers.StartImportResponseObject, error) {
	_, err := h.startImportHandler.Handle(ctx)
	if err != nil {
		return servers.StartImport500Response{}, nil
	}

	return servers.StartImport202JSONResponse{
		FailedRows: 0,
		SavedRows:  0,
		TotalRows:  0,
	}, nil
}
