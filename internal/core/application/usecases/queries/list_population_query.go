package queries

import (
	"context"
	"dataset-collections/internal/core/domain/model/kernel"
	"dataset-collections/internal/core/ports"
)

type ListPopulationQuery struct {
	CountryCode *kernel.CountryCode
	Year        *kernel.Year
	Limit       int
	Offset      int
}

type ListPopulationResult struct {
	Entries []kernel.PopulationEntry
	Total   int
}

type ListPopulationQueryHandler interface {
	Handle(ctx context.Context, query ListPopulationQuery) (ListPopulationResult, error)
}

type listPopulationQueryHandler struct {
	populationRepository ports.PopulationRepository
}

func NewListPopulationQueryHandler(repo ports.PopulationRepository) ListPopulationQueryHandler {
	return &listPopulationQueryHandler{
		populationRepository: repo,
	}
}

func (h *listPopulationQueryHandler) Handle(ctx context.Context, query ListPopulationQuery) (ListPopulationResult, error) {
	var entries []kernel.PopulationEntry
	var err error

	// Если указан код страны, получаем данные по нему
	if query.CountryCode != nil {
		entries, err = h.populationRepository.GetByCountryCode(ctx, *query.CountryCode)
		if err != nil {
			return ListPopulationResult{}, err
		}
	} else {
		// Иначе получаем все данные
		entries, err = h.populationRepository.GetAll(ctx)
		if err != nil {
			return ListPopulationResult{}, err
		}
	}

	// Применяем пагинацию
	total := len(entries)
	start := query.Offset
	end := start + query.Limit
	if end > total {
		end = total
	}
	if start > total {
		start = total
	}

	// Возвращаем срезанные данные
	return ListPopulationResult{
		Entries: entries[start:end],
		Total:   total,
	}, nil
}
