package postgres

import (
	"context"
	"dataset-collections/internal/core/domain/model/kernel"
	"dataset-collections/internal/core/ports"
	"dataset-collections/internal/pkg/errs"
)

// PopulationSaver сохраняет данные о населении в базу данных
type PopulationSaver struct {
	unitOfWork ports.UnitOfWork
}

// NewPopulationSaver создает новый экземпляр PopulationSaver
func NewPopulationSaver(uow ports.UnitOfWork) *PopulationSaver {
	return &PopulationSaver{
		unitOfWork: uow,
	}
}

// Save сохраняет массив PopulationEntry в базу данных
func (s *PopulationSaver) Save(ctx context.Context, entries []kernel.PopulationEntry) error {
	if len(entries) == 0 {
		return nil
	}

	// Начинаем транзакцию
	s.unitOfWork.Begin(ctx)
	defer func() {
		if err := s.unitOfWork.Rollback(); err != nil {
			// Логируем ошибку отката, но не возвращаем её
			// так как основная ошибка уже возвращена
		}
	}()

	// Получаем репозиторий через UnitOfWork
	repo := s.unitOfWork.PopulationRepository()

	// Сохраняем данные
	if err := repo.SaveAll(ctx, entries); err != nil {
		return errs.WrapInfrastructureError("failed to save population entries", err)
	}

	// Подтверждаем транзакцию
	if err := s.unitOfWork.Commit(ctx); err != nil {
		return errs.WrapInfrastructureError("failed to commit population save transaction", err)
	}

	return nil
}
