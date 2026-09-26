package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// we must return a pointer to the cache here,
// if we return the value, we get copies of the mutex
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]
	if !ok {
		return []byte{}, ok
	}
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.cacheEntries {
		if entry.isExpired(c.interval) {
			delete(c.cacheEntries, key)
		}
	}
}

// we don't need to check UTC explicitly here
// time.Since(t) is shorthand for time.Now().Sub(t)
func (e cacheEntry) isExpired(interval time.Duration) bool {
	return time.Since(e.createdAt) > interval
}
