package importjobrepo

import (
	"time"
)

// ImportJobDTO - структура для работы с базой данных
type ImportJobDTO struct {
	ID         string    `gorm:"primaryKey;type:uuid"`
	Status     string    `gorm:"not null"`
	StartedAt  time.Time `gorm:"not null"`
	FinishedAt *time.Time
	TotalRows  int    `gorm:"default:0"`
	SavedRows  int    `gorm:"default:0"`
	FailedRows int    `gorm:"default:0"`
	DurationMS int    `gorm:"default:0"`
	Error      string `gorm:"type:text"`
}

func (ImportJobDTO) TableName() string {
	return "import_job_dtos"
}
