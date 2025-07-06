package ports

import (
	"context"
	"io"
)

// Fetcher downloads data from a given source URL.
type Fetcher interface {
	Fetch(ctx context.Context) (io.Reader, error)
}
