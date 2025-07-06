package ports

import (
	"context"
	"dataset-collections/internal/core/domain/model/importjob"
)

type ImportJobRepository interface {
	Save(ctx context.Context, job importjob.ImportJob) error
	GetByID(ctx context.Context, jobID string) (*importjob.ImportJob, error)
	Update(ctx context.Context, job importjob.ImportJob) error
}
