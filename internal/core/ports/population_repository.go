package ports

import (
	"context"
	"dataset-collections/internal/core/domain/model/kernel"
)

type PopulationRepository interface {
	SaveAll(ctx context.Context, entries []kernel.PopulationEntry) error
	GetAll(ctx context.Context) ([]kernel.PopulationEntry, error)
	GetByCountryCode(ctx context.Context, countryCode kernel.CountryCode) ([]kernel.PopulationEntry, error)
}
