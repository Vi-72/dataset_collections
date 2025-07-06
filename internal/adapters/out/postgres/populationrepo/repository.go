package populationrepo

import (
	"context"
	"task-processing-service/internal/core/domain/model/kernel"
	"task-processing-service/internal/core/ports"
	"task-processing-service/internal/pkg/errs"

	"gorm.io/gorm"
)

var _ ports.PopulationRepository = &Repository{}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveAll(ctx context.Context, entries []kernel.PopulationEntry) error {
	dtos := make([]PopulationDTO, len(entries))
	for i, e := range entries {
		dtos[i] = DomainToDTO(e)
	}

	if err := r.db.WithContext(ctx).Create(&dtos).Error; err != nil {
		return errs.WrapInfrastructureError("failed to save population entries", err)
	}
	return nil
}
