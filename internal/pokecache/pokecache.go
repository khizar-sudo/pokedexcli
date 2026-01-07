package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	store map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		store: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if val, ok := cache.store[key]; !ok {
		return []byte{}, false
	} else {
		return val.val, true
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		now := time.Now()
		for key, val := range cache.store {
			if now.Sub(val.createdAt) > interval {
				cache.mu.Lock()
				delete(cache.store, key)
				cache.mu.Unlock()
			}
		}
	}
}
