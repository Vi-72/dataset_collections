package cmd

import (
	"task-processing-service/internal/adapters/in/http"
	memrepo "task-processing-service/internal/adapters/out/memory/taskrepo"
	"task-processing-service/internal/adapters/out/postgres/populationrepo"
	"task-processing-service/internal/core/application/usecases/commands"
	"task-processing-service/internal/core/application/usecases/queries"
	importersvc "task-processing-service/internal/core/domain/services/importer"
	tasksvc "task-processing-service/internal/core/domain/services/task"
	taskports "task-processing-service/internal/core/ports"
	"task-processing-service/internal/generated/servers"
)

type CompositionRoot struct {
	configs        Config
	taskRunner     *tasksvc.Runner
	taskRepository taskports.TaskRepository

	closers []Closer
}

func NewCompositionRoot(configs Config) *CompositionRoot {
	repo := memrepo.NewInMemoryTaskRepository()

	return &CompositionRoot{
		taskRepository: repo,
		configs:        configs,
	}
}

func (cr *CompositionRoot) NewTaskRunner() *tasksvc.Runner {
	if cr.taskRunner == nil {
		cr.taskRunner = tasksvc.NewTaskRunner(
			cr.configs.TaskRunnerWorkers,
			cr.configs.TaskRunnerQueueSize,
		)
	}
	return cr.taskRunner
}

func (cr *CompositionRoot) NewExecutionService() taskports.TaskExecutionService {
	return tasksvc.NewExecutionService(cr.NewTaskRunner())
}

func (cr *CompositionRoot) TaskRepository() taskports.TaskRepository {
	return cr.taskRepository
}

func (cr *CompositionRoot) NewExecuteTaskCommandHandler() commands.ExecuteTaskCommandHandler {
	handler, _ := commands.NewExecuteTaskCommandHandler(cr.NewExecutionService(), cr.taskRepository)
	return handler
}

func (cr *CompositionRoot) NewGetTaskQueryHandler() queries.GetTaskQueryHandler {
	return queries.NewGetTaskQueryHandler(cr.TaskRepository())
}

func (cr *CompositionRoot) NewDeleteTaskCommandHandler() queries.DeleteTaskQueryHandler {
	return queries.NewDeleteTaskQueryHandler(cr.taskRepository)
}

func (cr *CompositionRoot) NewTaskHandler() servers.StrictServerInterface {
	return http.NewTaskHandler(
		cr.NewExecuteTaskCommandHandler(),
		cr.NewGetTaskQueryHandler(),
		cr.NewDeleteTaskCommandHandler(),
	)
}

func (cr *CompositionRoot) NewImporterService() importersvc.Service {
	return importersvc.NewImporterService()
}

func (cr *CompositionRoot) PopulationRepository() ports.PopulationRepository {
	if cr.populationRepo == nil {
		cr.populationRepo = populationrepo.NewRepository(cr.db) // cr.db — *gorm.DB
	}
	return cr.populationRepo
}
