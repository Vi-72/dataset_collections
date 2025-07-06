package importjobrepo

import (
	"context"
	"dataset-collections/internal/core/domain/model/importjob"
	"dataset-collections/internal/core/ports"
	"dataset-collections/internal/pkg/errs"
	"gorm.io/gorm"
)

var _ ports.ImportJobRepository = &Repository{}

type Repository struct {
	tracker ports.Tracker
}

func NewRepository(tracker ports.Tracker) (*Repository, error) {
	if tracker == nil {
		return nil, errs.NewValueIsRequiredError("tracker")
	}
	return &Repository{tracker: tracker}, nil
}

func (r *Repository) Save(ctx context.Context, job importjob.ImportJob) error {
	dto := DomainToDTO(job)

	isInTransaction := r.tracker.InTx()
	if !isInTransaction {
		r.tracker.Begin(ctx)
	}
	tx := r.tracker.Tx()

	if err := tx.WithContext(ctx).Create(&dto).Error; err != nil {
		if !isInTransaction {
			_ = r.tracker.Rollback()
		}
		return errs.WrapInfrastructureError("failed to save import job", err)
	}

	if !isInTransaction {
		if err := r.tracker.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, jobID string) (*importjob.ImportJob, error) {
	var dto ImportJobDTO

	db := r.tracker.Db()
	if err := db.WithContext(ctx).Where("id = ?", jobID).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewObjectNotFoundError("import job not found", jobID)
		}
		return nil, errs.WrapInfrastructureError("failed to get import job", err)
	}

	job, err := DtoToDomain(dto)
	if err != nil {
		return nil, errs.WrapInfrastructureError("failed to convert dto to domain", err)
	}

	return &job, nil
}

func (r *Repository) Update(ctx context.Context, job importjob.ImportJob) error {
	dto := DomainToDTO(job)

	isInTransaction := r.tracker.InTx()
	if !isInTransaction {
		r.tracker.Begin(ctx)
	}
	tx := r.tracker.Tx()

	if err := tx.WithContext(ctx).Save(&dto).Error; err != nil {
		if !isInTransaction {
			_ = r.tracker.Rollback()
		}
		return errs.WrapInfrastructureError("failed to update import job", err)
	}

	if !isInTransaction {
		if err := r.tracker.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}
