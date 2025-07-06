package populationrepo

import "task-processing-service/internal/core/domain/model/kernel"

func DomainToDTO(entry kernel.PopulationEntry) PopulationDTO {
	return PopulationDTO{
		CountryName: entry.CountryName(),
		CountryCode: entry.CountryCode().Value(),
		Year:        entry.Year().Value(),
		Population:  entry.Population(),
	}
}

func DtoToDomain(dto PopulationDTO) (kernel.PopulationEntry, error) {
	code, err := kernel.NewCountryCode(dto.CountryCode)
	if err != nil {
		return kernel.PopulationEntry{}, err
	}

	year, err := kernel.NewYear(dto.Year)
	if err != nil {
		return kernel.PopulationEntry{}, err
	}

	return kernel.NewPopulationEntry(dto.CountryName, code, year, dto.Population)
}
