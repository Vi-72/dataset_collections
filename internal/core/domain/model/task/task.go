package task

import (
	"github.com/google/uuid"
	"task-processing-service/internal/pkg/ddd"
	"task-processing-service/internal/pkg/errs"
	"time"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "in_progress"
	StatusCompleted Status = "done"
	StatusFailed    Status = "failed"
)

type Task struct {
	*ddd.BaseAggregate[uuid.UUID]
	status    Status
	createdAt time.Time
	duration  *time.Duration
	result    *string
	errorMsg  *string
}

func NewTask(id uuid.UUID, now time.Time) (*Task, error) {
	if id == uuid.Nil {
		return nil, errs.NewValueIsRequiredError("TaskID")
	}

	return &Task{
		BaseAggregate: ddd.NewBaseAggregate(id),
		status:        StatusPending,
		createdAt:     now,
	}, nil
}

func (t *Task) MarkRunning() {
	t.status = StatusRunning
}

func (t *Task) Complete(result string, duration time.Duration) {
	t.status = StatusCompleted
	t.result = &result
	t.duration = &duration
	t.BaseAggregate.RaiseDomainEvent(NewCompletedDomainEvent(t))
}

func (t *Task) Fail(errMsg string, duration time.Duration) {
	t.status = StatusFailed
	t.errorMsg = &errMsg
	t.duration = &duration
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) Duration() *time.Duration {
	return t.duration
}

func (t *Task) Result() *string {
	return t.result
}

func (t *Task) ErrorMessage() *string {
	return t.errorMsg
}

func RestoreTask(
	id uuid.UUID,
	status Status,
	createdAt time.Time,
	duration *time.Duration,
	result *string,
	errorMsg *string,
) *Task {
	return &Task{
		BaseAggregate: ddd.NewBaseAggregate(id),
		status:        status,
		createdAt:     createdAt,
		duration:      duration,
		result:        result,
		errorMsg:      errorMsg,
	}
}
