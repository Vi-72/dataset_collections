package ports

import (
	"context"
)

type UnitOfWork interface {
	Begin(ctx context.Context)
	Commit(ctx context.Context) error
	Rollback() error
	PopulationRepository() PopulationRepository
	ImportJobRepository() ImportJobRepository
}
