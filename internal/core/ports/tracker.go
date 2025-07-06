package ports

import (
	"context"
	"gorm.io/gorm"
	"task-processing-service/internal/pkg/ddd"
)

type Tracker interface {
	Tx() *gorm.DB
	Db() *gorm.DB
	InTx() bool
	Track(agg ddd.AggregateRoot)
	Begin(ctx context.Context)
	Commit(ctx context.Context) error
}
