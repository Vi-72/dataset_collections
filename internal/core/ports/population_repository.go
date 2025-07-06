package ports

import (
	"context"
	"task-processing-service/internal/core/domain/model/kernel"
)

type PopulationRepository interface {
	SaveAll(ctx context.Context, entries []kernel.PopulationEntry) error
}
