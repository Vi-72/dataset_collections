package parser

import (
	"bufio"
	"dataset-collections/internal/core/domain/model/kernel"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// DataHubCSVParser парсит CSV данные с datahub.io
type DataHubCSVParser struct{}

// NewDataHubCSVParser создает новый экземпляр парсера
func NewDataHubCSVParser() *DataHubCSVParser {
	return &DataHubCSVParser{}
}

// Parse парсит CSV данные в массив PopulationEntry
func (p *DataHubCSVParser) Parse(r io.Reader) ([]kernel.PopulationEntry, error) {
	reader := csv.NewReader(bufio.NewReader(r))
	reader.TrimLeadingSpace = true

	// Читаем заголовки
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV headers: %w", err)
	}

	// Создаем карту индексов колонок
	colMap := make(map[string]int)
	for i, header := range headers {
		colMap[strings.ToLower(strings.TrimSpace(header))] = i
	}

	// Проверяем наличие обязательных колонок
	requiredColumns := []string{"country name", "country code", "year", "value"}
	for _, col := range requiredColumns {
		if _, exists := colMap[col]; !exists {
			return nil, fmt.Errorf("missing required column: %s", col)
		}
	}

	var entries []kernel.PopulationEntry

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("CSV parse error: %v", err)
			continue
		}

		entry, err := p.parseRow(colMap, record)
		if err != nil {
			log.Printf("CSV row parse error: %v", err)
			continue
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

// parseRow парсит одну строку CSV в PopulationEntry
func (p *DataHubCSVParser) parseRow(colMap map[string]int, row []string) (kernel.PopulationEntry, error) {
	// Получаем значения из колонок
	countryName := strings.TrimSpace(row[colMap["country name"]])
	countryCodeStr := strings.TrimSpace(row[colMap["country code"]])
	yearStr := strings.TrimSpace(row[colMap["year"]])
	valueStr := strings.TrimSpace(row[colMap["value"]])

	// Парсим год
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return kernel.PopulationEntry{}, fmt.Errorf("invalid year '%s': %w", yearStr, err)
	}

	// Парсим население
	population, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return kernel.PopulationEntry{}, fmt.Errorf("invalid population value '%s': %w", valueStr, err)
	}

	// Создаем доменные объекты
	countryCode, err := kernel.NewCountryCode(countryCodeStr)
	if err != nil {
		return kernel.PopulationEntry{}, fmt.Errorf("invalid country code '%s': %w", countryCodeStr, err)
	}

	yearObj, err := kernel.NewYear(year)
	if err != nil {
		return kernel.PopulationEntry{}, fmt.Errorf("invalid year %d: %w", year, err)
	}

	// Создаем PopulationEntry
	entry, err := kernel.NewPopulationEntry(countryName, countryCode, yearObj, population)
	if err != nil {
		return kernel.PopulationEntry{}, fmt.Errorf("failed to create population entry: %w", err)
	}

	return entry, nil
}
