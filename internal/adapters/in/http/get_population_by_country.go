package http

import (
	"context"
	"dataset-collections/internal/core/application/usecases/queries"
	"dataset-collections/internal/core/domain/model/kernel"
	"dataset-collections/internal/generated/servers"
)

func (h *ApiHandler) GetPopulationByCountry(ctx context.Context, request servers.GetPopulationByCountryRequestObject) (servers.GetPopulationByCountryResponseObject, error) {
	// Создаем CountryCode из строки
	countryCode, err := kernel.NewCountryCode(request.CountryCode)
	if err != nil {
		return servers.GetPopulationByCountry400JSONResponse{
			Error: "invalid country code: " + err.Error(),
		}, nil
	}

	query := queries.ListPopulationQuery{
		CountryCode: &countryCode,
		Limit:       100,
		Offset:      0,
	}
	
	result, err := h.listPopulationHandler.Handle(ctx, query)
	if err != nil {
		return servers.GetPopulationByCountry500JSONResponse{
			Error: err.Error(),
		}, nil
	}

	// Преобразуем доменные объекты в DTO для ответа
	var response []servers.PopulationEntry
	for _, entry := range result.Entries {
		response = append(response, servers.PopulationEntry{
			CountryName: entry.CountryName(),
			CountryCode: entry.CountryCode().Value(),
			Year:        entry.Year().Value(),
			Value:       int(entry.Population()),
		})
	}

	return servers.GetPopulationByCountry200JSONResponse(response), nil
} 