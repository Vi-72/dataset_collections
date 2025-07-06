package http

import (
	"task-processing-service/internal/core/application/usecases/commands"
	"task-processing-service/internal/core/application/usecases/queries"
)

type TaskHandler struct {
	executeHandler commands.ExecuteTaskCommandHandler
	getTaskQuery   queries.GetTaskQueryHandler
	deleteHandler  queries.DeleteTaskQueryHandler
}

func NewTaskHandler(
	executeHandler commands.ExecuteTaskCommandHandler,
	getTaskQuery queries.GetTaskQueryHandler,
	deleteHandler queries.DeleteTaskQueryHandler,
) *TaskHandler {
	return &TaskHandler{
		executeHandler: executeHandler,
		getTaskQuery:   getTaskQuery,
		deleteHandler:  deleteHandler,
	}
}
