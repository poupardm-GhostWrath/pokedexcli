package pokeapi

import (
	"net/http"
	"time"
	"github.com/poupardm-GhostWrath/pokedexcli/internal/pokecache"
)

// Client
type Client struct {
	httpClient		http.Client
	pokeCache		pokecache.Cache
}

// NewClient
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache:	pokecache.NewCache(cacheInterval),
	}
}