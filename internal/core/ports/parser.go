package ports

import (
	"dataset-collections/internal/core/domain/model/kernel"
	"io"
)

// Parser parses population data from a stream.
type Parser interface {
	Parse(r io.Reader) ([]kernel.PopulationEntry, error)
}
