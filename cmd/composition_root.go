package cmd

import (
	"dataset-collections/internal/adapters/in/http"
	"dataset-collections/internal/adapters/out/fetcher"
	"dataset-collections/internal/adapters/out/parser"
	"dataset-collections/internal/adapters/out/postgres"
	"dataset-collections/internal/core/application/usecases/commands"
	"dataset-collections/internal/core/application/usecases/queries"
	importersvc "dataset-collections/internal/core/domain/services/importer"
	"dataset-collections/internal/core/ports"
	"dataset-collections/internal/generated/servers"
	"gorm.io/gorm"
	"log"
)

type CompositionRoot struct {
	configs Config
	db      *gorm.DB

	closers []Closer
}

func NewCompositionRoot(configs Config, db *gorm.DB) *CompositionRoot {
	return &CompositionRoot{
		configs: configs,
		db:      db,
	}
}

func (cr *CompositionRoot) NewUnitOfWork() ports.UnitOfWork {
	unitOfWork, err := postgres.NewUnitOfWork(cr.db)
	if err != nil {
		log.Fatalf("cannot create UnitOfWork: %v", err)
	}
	return unitOfWork
}

func (cr *CompositionRoot) PopulationRepository() ports.PopulationRepository {
	return cr.NewUnitOfWork().PopulationRepository()
}

func (cr *CompositionRoot) NewListPopulationQueryHandler() queries.ListPopulationQueryHandler {
	return queries.NewListPopulationQueryHandler(cr.PopulationRepository())
}

func (cr *CompositionRoot) NewGetImportJobStatusQueryHandler() queries.GetImportJobStatusQueryHandler {
	return queries.NewGetImportJobStatusQueryHandler(cr.NewUnitOfWork().ImportJobRepository())
}

func (cr *CompositionRoot) NewApiHandler() servers.StrictServerInterface {
	handlers, err := http.NewApiHandler(
		cr.NewStartImportCommandHandler(),
		cr.NewListPopulationQueryHandler(),
	)

	if err != nil {
		log.Fatalf("Ошибка инициализации HTTP Server: %v", err)
	}

	return handlers
}

func (cr *CompositionRoot) NewImporterService() importersvc.Service {
	hubFetcher := fetcher.NewDataHubFetcher()
	hubCSVParser := parser.NewDataHubCSVParser()
	saver := postgres.NewPopulationSaver(cr.NewUnitOfWork())

	return importersvc.NewService(hubFetcher, hubCSVParser, saver)
}

func (cr *CompositionRoot) NewStartImportCommandHandler() commands.StartImportCommandHandler {
	return commands.NewStartImportCommandHandler(cr.NewUnitOfWork(), cr.NewImporterService(), cr.configs.PopulationCsvURL)
}
