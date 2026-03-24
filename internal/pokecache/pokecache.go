package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cacheEntries	map[string]cacheEntry
	mux				*sync.RWMutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) Cache {
	entries := make(map[string]cacheEntry)
	mu := &sync.RWMutex{}
	cache := Cache{
		cacheEntries: 	entries,
		mux:			mu,
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
		}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, ok := c.cacheEntries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.cacheEntries {
			if entry.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.cacheEntries, key)
			}
		}
		c.mux.Unlock()
	}
}