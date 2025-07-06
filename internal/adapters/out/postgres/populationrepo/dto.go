package populationrepo

type PopulationDTO struct {
	CountryName string `gorm:"primaryKey"`
	CountryCode string `gorm:"primaryKey"`
	Year        int    `gorm:"primaryKey"`
	Population  int64
}

func (PopulationDTO) TableName() string {
	return "population"
}
