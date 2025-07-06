package importer

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"task-processing-service/internal/core/domain/model/kernel"
)

var ErrInvalidRow = errors.New("invalid CSV row: expected 4 fields")

type Service interface {
	ParseCSV(r io.Reader) ([]kernel.PopulationEntry, error)
}

type importerService struct{}

func NewImporterService() Service {
	return &importerService{}
}

func (s *importerService) ParseCSV(r io.Reader) ([]kernel.PopulationEntry, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true

	// Чтение заголовков
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var entries []kernel.PopulationEntry

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		entry, err := parseRow(record)
		if err != nil {
			continue // можно логировать
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func parseRow(fields []string) (kernel.PopulationEntry, error) {
	if len(fields) != 4 {
		return kernel.PopulationEntry{}, ErrInvalidRow
	}

	yearInt, err := strconv.Atoi(fields[2])
	if err != nil {
		return kernel.PopulationEntry{}, err
	}

	populationInt, err := strconv.ParseInt(fields[3], 10, 64)
	if err != nil {
		return kernel.PopulationEntry{}, err
	}

	code, err := kernel.NewCountryCode(fields[1])
	if err != nil {
		return kernel.PopulationEntry{}, err
	}

	year, err := kernel.NewYear(yearInt)
	if err != nil {
		return kernel.PopulationEntry{}, err
	}

	return kernel.NewPopulationEntry(fields[0], code, year, populationInt)
}
