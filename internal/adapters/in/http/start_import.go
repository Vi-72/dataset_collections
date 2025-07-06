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
	
	result, err := h.startImportHandler.Handle(ctx, cmd)
	if err != nil {
		return servers.StartImport500JSONResponse{
			Error: "failed to start import: " + err.Error(),
		}, nil
	}

	return servers.StartImport202JSONResponse{
		JobId: result.JobID,
	}, nil
}
