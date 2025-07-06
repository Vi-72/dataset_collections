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

	if err := tx.WithContext(ctx).Save(&dtos).Error; err != nil {
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

func (r *Repository) GetAll(ctx context.Context) ([]kernel.PopulationEntry, error) {
	var dtos []PopulationDTO

	db := r.tracker.Db()
	if err := db.WithContext(ctx).Find(&dtos).Error; err != nil {
		return nil, errs.WrapInfrastructureError("failed to get population entries", err)
	}

	entries := make([]kernel.PopulationEntry, len(dtos))
	for i, dto := range dtos {
		entry, err := DtoToDomain(dto)
		if err != nil {
			return nil, errs.WrapInfrastructureError("failed to convert dto to domain", err)
		}
		entries[i] = entry
	}

	return entries, nil
}
