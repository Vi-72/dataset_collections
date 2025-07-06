package kernel

import (
	"fmt"
)

type PopulationEntry struct {
	countryName string
	countryCode CountryCode
	year        Year
	population  int64
}

func NewPopulationEntry(countryName string, countryCode CountryCode, year Year, population int64) (PopulationEntry, error) {
	if countryName == "" {
		return PopulationEntry{}, fmt.Errorf("country name cannot be empty")
	}
	if population < 0 {
		return PopulationEntry{}, fmt.Errorf("population must be non-negative: %d", population)
	}
	return PopulationEntry{
		countryName: countryName,
		countryCode: countryCode,
		year:        year,
		population:  population,
	}, nil
}

func (p PopulationEntry) CountryName() string {
	return p.countryName
}

func (p PopulationEntry) CountryCode() CountryCode {
	return p.countryCode
}

func (p PopulationEntry) Year() Year {
	return p.year
}

func (p PopulationEntry) Population() int64 {
	return p.population
}

func (p PopulationEntry) Equals(other PopulationEntry) bool {
	return p.countryCode.Equals(other.countryCode) &&
		p.year.Equals(other.year) &&
		p.population == other.population &&
		p.countryName == other.countryName
}
