package importjob

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type ImportResult struct {
	TotalRows  int
	SavedRows  int
	FailedRows int
	DurationMS int
	Error      string
}

type ImportJob struct {
	ID         uuid.UUID
	Status     Status
	StartedAt  time.Time
	FinishedAt *time.Time
	Result     *ImportResult
}

func NewImportJob() ImportJob {
	return ImportJob{
		ID:        uuid.New(),
		Status:    StatusPending,
		StartedAt: time.Now(),
	}
}

func (j *ImportJob) MarkInProgress() {
	j.Status = StatusInProgress
}

func (j *ImportJob) MarkCompleted(result ImportResult) {
	now := time.Now()
	j.Status = StatusCompleted
	j.FinishedAt = &now
	j.Result = &result
}

func (j *ImportJob) MarkFailed(err error) {
	now := time.Now()
	j.Status = StatusFailed
	j.FinishedAt = &now
	j.Result = &ImportResult{
		Error: err.Error(),
	}
}
