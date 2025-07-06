package ports

import (
	"context"
	"dataset-collections/internal/pkg/ddd"
	"gorm.io/gorm"
)

type Tracker interface {
	Tx() *gorm.DB
	Db() *gorm.DB
	InTx() bool
	Track(agg ddd.AggregateRoot)
	Begin(ctx context.Context)
	Commit(ctx context.Context) error
	Rollback() error
}
