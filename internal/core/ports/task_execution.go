package ports

import (
	"context"
	"task-processing-service/internal/core/domain/model/task"
)

type TaskExecutionService interface {
	Execute(ctx context.Context, t *task.Task) error
}
