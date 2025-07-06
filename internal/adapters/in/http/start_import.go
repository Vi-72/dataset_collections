package http

import (
	"context"
	"dataset-collections/internal/core/application/usecases/commands"
	"dataset-collections/internal/generated/servers"
)

func (h *ApiHandler) StartImport(ctx context.Context, request servers.StartImportRequestObject) (servers.StartImportResponseObject, error) {
	// Создаем команду с дефолтным URL (так как StartImportRequestObject пустая)
	cmd := commands.StartImportCommand{
		SourceURL: "", // Будет использован дефолтный URL из конфига
	}

	_, err := h.startImportHandler.Handle(ctx, cmd)
	if err != nil {
		return servers.StartImport500Response{}, nil
	}

	return servers.StartImport202JSONResponse{
		FailedRows: 0,
		SavedRows:  0,
		TotalRows:  0,
	}, nil
}
