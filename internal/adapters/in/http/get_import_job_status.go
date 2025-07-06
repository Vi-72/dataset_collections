package http

import (
	"context"
	"dataset-collections/internal/core/application/usecases/queries"
	"dataset-collections/internal/generated/servers"
	"time"
)

func (h *ApiHandler) GetImportJobStatus(ctx context.Context, request servers.GetImportJobStatusRequestObject) (servers.GetImportJobStatusResponseObject, error) {
	// Получаем статус джобы
	status, err := h.getImportJobStatusHandler.Handle(ctx, queries.GetImportJobStatusQuery{
		JobID: request.JobId,
	})
	if err != nil {
		// Проверяем тип ошибки для правильного HTTP кода
		if err.Error() == "import job not found" {
			return servers.GetImportJobStatus404Response{}, nil
		}
		return servers.GetImportJobStatus500Response{}, nil
	}

	// Конвертируем время в правильный формат
	startedAt := status.StartedAt
	var finishedAt *time.Time
	if status.FinishedAt != nil {
		finishedAt = status.FinishedAt
	}

	// Конвертируем в указатели для опциональных полей
	totalRows := &status.TotalRows
	savedRows := &status.SavedRows
	failedRows := &status.FailedRows
	durationMs := &status.DurationMS
	var errorPtr *string
	if status.Error != "" {
		errorPtr = &status.Error
	}

	return servers.GetImportJobStatus200JSONResponse{
		JobId:      status.JobID,
		Status:     servers.ImportJobStatusStatus(status.Status),
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
		TotalRows:  totalRows,
		SavedRows:  savedRows,
		FailedRows: failedRows,
		DurationMs: durationMs,
		Error:      errorPtr,
	}, nil
}
