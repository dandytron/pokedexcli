package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{
		val:       value,
		createdAt: time.Now().UTC(),
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	return entry.val, ok
}

// reapLoop -
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// this will run every interval - if 5 mins, every 5 mins
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	intervalAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(intervalAgo) {
			delete(c.cache, k)
		}
	}
}
