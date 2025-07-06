package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Fetcher загружает CSV-данные по URL
// Используется для загрузки данных из внешних источников, таких как datahub.io
type Fetcher interface {
	Fetch(ctx context.Context, url string) (io.Reader, error)
}

// HTTPFetcher реализует Fetcher через стандартную библиотеку http
type HTTPFetcher struct {
	client     *http.Client
	defaultURL string
}

// NewHTTPFetcher создаёт новый HTTPFetcher с заданным базовым URL (по умолчанию, если не передан URL в Fetch)
func NewHTTPFetcher(defaultURL string) *HTTPFetcher {
	return &HTTPFetcher{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		defaultURL: defaultURL,
	}
}

// Fetch загружает CSV-файл по заданному URL или использует URL по умолчанию
func (f *HTTPFetcher) Fetch(ctx context.Context, url string) (io.Reader, error) {
	if url == "" {
		url = f.defaultURL
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "dataset-collections/1.0")
	req.Header.Set("Accept", "text/csv")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code %d from URL %s", resp.StatusCode, url)
	}

	return resp.Body, nil
}
