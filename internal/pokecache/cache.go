package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
	return c
}
