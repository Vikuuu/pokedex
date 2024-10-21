package pokecache

import (
	"sync"
	"time"
)

// structs

type Cache struct {
	mu   *sync.Mutex
	data map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// New Cache

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mu:   &sync.Mutex{},
		data: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.data[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(d time.Duration) {
	for {
		time.Sleep(d)
		c.mu.Lock()
		for key, cache := range c.data {
			if time.Since(cache.createdAt) > d {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
