package ports

import (
	"context"
	"dataset-collections/internal/core/domain/model/kernel"
)

type PopulationRepository interface {
	SaveAll(ctx context.Context, entries []kernel.PopulationEntry) error
}
