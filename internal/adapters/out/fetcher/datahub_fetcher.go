package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DataHubFetcher загружает данные с datahub.io API
type DataHubFetcher struct {
	client  *http.Client
	baseURL string
}

// NewDataHubFetcher создает новый экземпляр DataHubFetcher
func NewDataHubFetcher() *DataHubFetcher {
	return &DataHubFetcher{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: "https://datahub.io/core/population",
	}
}

// Fetch загружает CSV файл с данными о населении
func (f *DataHubFetcher) Fetch(ctx context.Context, url string) (io.Reader, error) {
	// Если URL не указан, используем стандартный URL datahub.io
	if url == "" {
		url = f.baseURL + "/_r/-/data/population.csv"
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Добавляем заголовки для корректной работы с API
	req.Header.Set("User-Agent", "dataset-collections/1.0")
	req.Header.Set("Accept", "text/csv")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d from URL %s", resp.StatusCode, url)
	}

	return resp.Body, nil
} 