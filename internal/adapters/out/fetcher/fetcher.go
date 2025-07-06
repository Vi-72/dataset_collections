package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

// CSVFetcher отвечает за загрузку CSV-файлов по URL
type CSVFetcher interface {
	Fetch(url string) (io.Reader, error)
}

type httpCSVFetcher struct{}

func NewHTTPFetcher() CSVFetcher {
	return &httpCSVFetcher{}
}

func (f *httpCSVFetcher) Fetch(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d from URL %s", resp.StatusCode, url)
	}
	return resp.Body, nil
}
