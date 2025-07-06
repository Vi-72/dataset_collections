package importjobrepo

import (
	"dataset-collections/internal/core/domain/model/importjob"
	"github.com/google/uuid"
)

// DomainToDTO преобразует доменный объект в DTO
func DomainToDTO(job importjob.ImportJob) ImportJobDTO {
	dto := ImportJobDTO{
		ID:         job.ID.String(),
		SourceURL:  job.SourceURL,
		Status:     string(job.Status),
		StartedAt:  job.StartedAt,
		FinishedAt: job.FinishedAt,
	}

	if job.Result != nil {
		dto.TotalRows = job.Result.TotalRows
		dto.SavedRows = job.Result.SavedRows
		dto.FailedRows = job.Result.FailedRows
		dto.DurationMS = job.Result.DurationMS
		dto.Error = job.Result.Error
	}

	return dto
}

// DtoToDomain преобразует DTO в доменный объект
func DtoToDomain(dto ImportJobDTO) (importjob.ImportJob, error) {
	jobID, err := uuid.Parse(dto.ID)
	if err != nil {
		return importjob.ImportJob{}, err
	}

	job := importjob.ImportJob{
		ID:        jobID,
		SourceURL: dto.SourceURL,
		Status:    importjob.Status(dto.Status),
		StartedAt: dto.StartedAt,
	}

	if dto.FinishedAt != nil {
		job.FinishedAt = dto.FinishedAt
	}

	if dto.TotalRows > 0 || dto.SavedRows > 0 || dto.FailedRows > 0 || dto.DurationMS > 0 || dto.Error != "" {
		job.Result = &importjob.ImportResult{
			TotalRows:  dto.TotalRows,
			SavedRows:  dto.SavedRows,
			FailedRows: dto.FailedRows,
			DurationMS: dto.DurationMS,
			Error:      dto.Error,
		}
	}

	return job, nil
} 