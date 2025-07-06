package taskrepo

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"task-processing-service/internal/core/domain/model/task"
	"task-processing-service/internal/pkg/errs"
)

type InMemoryTaskRepository struct {
	data map[uuid.UUID]*task.Task
	mu   sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		data: make(map[uuid.UUID]*task.Task),
	}
}

func (r *InMemoryTaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, exists := r.data[id]
	if !exists {
		return nil, errs.ErrObjectNotFound
	}
	return task, nil
}

func (r *InMemoryTaskRepository) GetAll(ctx context.Context) ([]*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := make([]*task.Task, 0, len(r.data))
	for _, task := range r.data {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) Save(ctx context.Context, task *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[task.ID()] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}
