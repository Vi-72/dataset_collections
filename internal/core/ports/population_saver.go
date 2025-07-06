package ports

import (
	"context"
	"dataset-collections/internal/core/domain/model/kernel"
)

// Saver persists population entries.
type Saver interface {
	Save(ctx context.Context, entries []kernel.PopulationEntry) error
}
