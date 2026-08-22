package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ce, ok := c.entries[key]
	if !ok {
		return make([]byte, 0), false
	}
	return ce.val, true
}

func (c *Cache) reapLoop() {
	go func() {
		ticker := time.NewTicker(c.interval)
		defer ticker.Stop()

		for range ticker.C {
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
