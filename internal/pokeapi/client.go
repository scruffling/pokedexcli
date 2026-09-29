package pokeapi

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/scruffling/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(cacheInterval),
	}
}

// StatusError is returned when the API responds with a non-200 status.
type StatusError struct {
	StatusCode int
	URL        string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("unexpected status %d from %s", e.StatusCode, e.URL)
}

// getBody returns the raw response bytes for url, using the cache when possible.
// Non-200 responses return a *StatusError and are not cached.
func (c *Client) getBody(url string) ([]byte, error) {
	if data, ok := c.pokeCache.Get(url); ok {
		slog.Debug("Using cached data...")
		return data, nil
	}

	slog.Debug("No cached data available. Making request...")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request generation error: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, &StatusError{StatusCode: res.StatusCode, URL: url}
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	c.pokeCache.Add(url, data)
	return data, nil
}
