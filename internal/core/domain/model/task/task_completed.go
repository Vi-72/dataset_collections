package task

import (
	"github.com/google/uuid"
	"reflect"
	"task-processing-service/internal/pkg/ddd"
)

type CompletedDomainEvent struct {
	// base
	ID   uuid.UUID
	Name string

	// payload
	TaskID     uuid.UUID
	TaskStatus string

	isSet bool
}

func (e CompletedDomainEvent) GetID() uuid.UUID {
	return e.ID
}

func (e CompletedDomainEvent) GetName() string {
	return e.Name
}

func (e CompletedDomainEvent) IsEmpty() bool {
	return !e.isSet
}

func NewCompletedDomainEvent(t *Task) ddd.DomainEvent {
	event := CompletedDomainEvent{
		ID:         uuid.New(),
		TaskID:     t.ID(),
		TaskStatus: string(t.Status()),
		isSet:      true,
	}
	event.Name = reflect.TypeOf(event).Name()
	return &event
}
