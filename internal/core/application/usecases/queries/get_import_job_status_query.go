package queries

import (
	"context"
	"dataset-collections/internal/core/domain/model/importjob"
	"dataset-collections/internal/core/ports"
	"time"
)

type GetImportJobStatusQuery struct {
	JobID string
}

type GetImportJobStatusResult struct {
	JobID      string
	Status     importjob.Status
	StartedAt  time.Time
	FinishedAt *time.Time
	TotalRows  int
	SavedRows  int
	FailedRows int
	DurationMS int
	Error      string
}

type GetImportJobStatusQueryHandler interface {
	Handle(ctx context.Context, query GetImportJobStatusQuery) (GetImportJobStatusResult, error)
}

type getImportJobStatusQueryHandler struct {
	importJobRepository ports.ImportJobRepository
}

func NewGetImportJobStatusQueryHandler(repo ports.ImportJobRepository) GetImportJobStatusQueryHandler {
	return &getImportJobStatusQueryHandler{
		importJobRepository: repo,
	}
}

func (h *getImportJobStatusQueryHandler) Handle(ctx context.Context, query GetImportJobStatusQuery) (GetImportJobStatusResult, error) {
	// Получаем джобу из репозитория
	job, err := h.importJobRepository.GetByID(ctx, query.JobID)
	if err != nil {
		return GetImportJobStatusResult{}, err
	}

	// Формируем результат
	result := GetImportJobStatusResult{
		JobID:     job.ID.String(),
		Status:    job.Status,
		StartedAt: job.StartedAt,
	}

	if job.FinishedAt != nil {
		result.FinishedAt = job.FinishedAt
	}

	if job.Result != nil {
		result.TotalRows = job.Result.TotalRows
		result.SavedRows = job.Result.SavedRows
		result.FailedRows = job.Result.FailedRows
		result.DurationMS = job.Result.DurationMS
		result.Error = job.Result.Error
	}

	return result, nil
}
