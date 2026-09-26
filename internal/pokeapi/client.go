package pokeapi

import (
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
