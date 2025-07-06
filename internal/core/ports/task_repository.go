package ports

import (
	"context"
	"github.com/google/uuid"
	"task-processing-service/internal/core/domain/model/task"
)

type TaskRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error)
	GetAll(ctx context.Context) ([]*task.Task, error)
	Save(ctx context.Context, task *task.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}
