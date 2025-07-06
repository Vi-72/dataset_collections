package populationrepo

import (
	"context"
	"dataset-collections/internal/core/domain/model/kernel"
	"dataset-collections/internal/core/ports"
	"dataset-collections/internal/pkg/errs"
)

var _ ports.PopulationRepository = &Repository{}

type Repository struct {
	tracker ports.Tracker
}

func NewRepository(tracker ports.Tracker) (*Repository, error) {
	if tracker == nil {
		return nil, errs.NewValueIsRequiredError("tracker")
	}
	return &Repository{tracker: tracker}, nil
}

func (r *Repository) SaveAll(ctx context.Context, entries []kernel.PopulationEntry) error {
	if len(entries) == 0 {
		return nil
	}

	isInTransaction := r.tracker.InTx()
	if !isInTransaction {
		r.tracker.Begin(ctx)
	}
	tx := r.tracker.Tx()

	dtos := make([]PopulationDTO, len(entries))
	for i, e := range entries {
		dtos[i] = DomainToDTO(e)
	}

	if err := tx.WithContext(ctx).Create(&dtos).Error; err != nil {
		if !isInTransaction {
			_ = r.tracker.Rollback()
		}
		return errs.WrapInfrastructureError("failed to save population entries", err)
	}

	if !isInTransaction {
		if err := r.tracker.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}
