package commands

import (
	"context"
	"dataset-collections/internal/core/domain/model/importjob"
	"dataset-collections/internal/core/domain/services/importer"
	"dataset-collections/internal/core/ports"
)

type StartImportCommand struct {
	SourceURL string
}

type StartImportResult struct {
	JobID string
}

type StartImportCommandHandler interface {
	Handle(ctx context.Context, command StartImportCommand) (StartImportResult, error)
}

type startImportCommandHandler struct {
	unitOfWork       ports.UnitOfWork
	importerService  importer.Service
	defaultSourceURL string
}

func NewStartImportCommandHandler(
	unitOfWork ports.UnitOfWork,
	importerSvc importer.Service,
	defaultSourceURL string,
) StartImportCommandHandler {
	return &startImportCommandHandler{
		unitOfWork:       unitOfWork,
		importerService:  importerSvc,
		defaultSourceURL: defaultSourceURL,
	}
}

func (h *startImportCommandHandler) Handle(ctx context.Context, command StartImportCommand) (StartImportResult, error) {
	// Используем URL из команды или дефолтный URL
	sourceURL := command.SourceURL
	if sourceURL == "" {
		sourceURL = h.defaultSourceURL
	}

	// Создаём новый импорт-джоб
	job := importjob.NewImportJob(sourceURL)

	// Сохраняем джоб в репозиторий
	if err := h.unitOfWork.ImportJobRepository().Save(ctx, job); err != nil {
		return StartImportResult{}, err
	}

	// Запускаем импорт в фоне
	go func() {
		_, err := h.importerService.Start(context.Background(), &job)
		if err != nil {
			// TODO: логировать ошибку импорта
		}

		// Обновляем статус джоба после импорта
		if updateErr := h.unitOfWork.ImportJobRepository().Update(context.Background(), job); updateErr != nil {
			// TODO: логировать ошибку обновления
		}
	}()

	return StartImportResult{
		JobID: job.ID.String(),
	}, nil
}
