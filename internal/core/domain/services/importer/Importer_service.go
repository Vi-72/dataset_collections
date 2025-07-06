package importer

import (
	"context"
	"dataset-collections/internal/core/domain/model/importjob"
	"dataset-collections/internal/core/domain/model/kernel"
	"dataset-collections/internal/core/ports"
	"errors"
	"fmt"
	"time"
)

// Service defines the interface for starting a population import
// by fetching, parsing and saving data.
type Service interface {
	Start(ctx context.Context, job *importjob.ImportJob) (*importjob.ImportResult, error)
}

type service struct {
	fetcher ports.Fetcher
	parser  ports.Parser
	saver   ports.Saver
}

func NewService(fetcher ports.Fetcher, parser ports.Parser, saver ports.Saver) Service {
	return &service{
		fetcher: fetcher,
		parser:  parser,
		saver:   saver,
	}
}

func (s *service) Start(ctx context.Context, job *importjob.ImportJob) (*importjob.ImportResult, error) {
	if job == nil {
		return nil, errors.New("import job is nil")
	}

	startTime := time.Now()

	// Mark job as in progress
	job.MarkInProgress()

	// Fetch data from external URL
	reader, err := s.fetcher.Fetch(ctx)
	if err != nil {
		job.MarkFailed(fmt.Errorf("failed to fetch data: %w", err))
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	// Parse the data
	entries, err := s.parser.Parse(reader)
	if err != nil {
		job.MarkFailed(fmt.Errorf("failed to parse data: %w", err))
		return nil, fmt.Errorf("failed to parse data: %w", err)
	}

	// Filter valid entries
	valid := s.filterValid(entries)

	// Save valid entries
	if err := s.saver.Save(ctx, valid); err != nil {
		job.MarkFailed(fmt.Errorf("failed to save data: %w", err))
		return nil, fmt.Errorf("failed to save data: %w", err)
	}

	duration := time.Since(startTime)
	result := &importjob.ImportResult{
		TotalRows:  len(entries),
		SavedRows:  len(valid),
		FailedRows: len(entries) - len(valid),
		DurationMS: int(duration.Milliseconds()),
	}

	job.MarkCompleted(*result)
	return result, nil
}

func (s *service) filterValid(entries []kernel.PopulationEntry) []kernel.PopulationEntry {
	var valid []kernel.PopulationEntry
	for _, e := range entries {
		if e.Population() >= 0 && e.Year().Value() >= 1800 {
			valid = append(valid, e)
		}
	}
	return valid
}
