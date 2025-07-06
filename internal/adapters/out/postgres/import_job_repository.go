package postgres

import (
	"context"
	"dataset-collections/internal/core/domain/model/importjob"
	"dataset-collections/internal/core/ports"
	"dataset-collections/internal/pkg/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ImportJobRepository struct {
	db *gorm.DB
}

func NewImportJobRepository(db *gorm.DB) ports.ImportJobRepository {
	return &ImportJobRepository{db: db}
}

func (r *ImportJobRepository) Save(ctx context.Context, job importjob.ImportJob) error {
	dto := ImportJobDTO{
		ID:         job.ID.String(),
		SourceURL:  job.SourceURL,
		Status:     string(job.Status),
		StartedAt:  job.StartedAt,
		FinishedAt: job.FinishedAt,
	}

	if job.Result != nil {
		dto.TotalRows = job.Result.TotalRows
		dto.SavedRows = job.Result.SavedRows
		dto.FailedRows = job.Result.FailedRows
		dto.DurationMS = job.Result.DurationMS
		dto.Error = job.Result.Error
	}

	if err := r.db.WithContext(ctx).Create(&dto).Error; err != nil {
		return errs.WrapInfrastructureError("failed to save import job", err)
	}

	return nil
}

func (r *ImportJobRepository) GetByID(ctx context.Context, jobID string) (*importjob.ImportJob, error) {
	var dto ImportJobDTO
	
	if err := r.db.WithContext(ctx).Where("id = ?", jobID).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewObjectNotFoundError("import job not found")
		}
		return nil, errs.WrapInfrastructureError("failed to get import job", err)
	}

	return dto.ToDomain(), nil
}

func (r *ImportJobRepository) Update(ctx context.Context, job importjob.ImportJob) error {
	dto := ImportJobDTO{
		ID:         job.ID.String(),
		SourceURL:  job.SourceURL,
		Status:     string(job.Status),
		StartedAt:  job.StartedAt,
		FinishedAt: job.FinishedAt,
	}

	if job.Result != nil {
		dto.TotalRows = job.Result.TotalRows
		dto.SavedRows = job.Result.SavedRows
		dto.FailedRows = job.Result.FailedRows
		dto.DurationMS = job.Result.DurationMS
		dto.Error = job.Result.Error
	}

	if err := r.db.WithContext(ctx).Save(&dto).Error; err != nil {
		return errs.WrapInfrastructureError("failed to update import job", err)
	}

	return nil
}

// ImportJobDTO - структура для работы с базой данных
type ImportJobDTO struct {
	ID         string     `gorm:"primaryKey;type:uuid"`
	SourceURL  string     `gorm:"not null"`
	Status     string     `gorm:"not null"`
	StartedAt  time.Time  `gorm:"not null"`
	FinishedAt *time.Time
	TotalRows  int    `gorm:"default:0"`
	SavedRows  int    `gorm:"default:0"`
	FailedRows int    `gorm:"default:0"`
	DurationMS int    `gorm:"default:0"`
	Error      string `gorm:"type:text"`
}

func (dto ImportJobDTO) ToDomain() *importjob.ImportJob {
	jobID, _ := uuid.Parse(dto.ID)
	
	job := &importjob.ImportJob{
		ID:        jobID,
		SourceURL: dto.SourceURL,
		Status:    importjob.Status(dto.Status),
		StartedAt: dto.StartedAt,
	}

	if dto.FinishedAt != nil {
		job.FinishedAt = dto.FinishedAt
	}

	if dto.TotalRows > 0 || dto.SavedRows > 0 || dto.FailedRows > 0 || dto.DurationMS > 0 || dto.Error != "" {
		job.Result = &importjob.ImportResult{
			TotalRows:  dto.TotalRows,
			SavedRows:  dto.SavedRows,
			FailedRows: dto.FailedRows,
			DurationMS: dto.DurationMS,
			Error:      dto.Error,
		}
	}

	return job
} 