package lru

import (
	"sync"
)

type Cache[K comparable, V any] struct {
	size  uint
	mu    sync.RWMutex
	inner map[K]*el[K, V]
	lru   *list[K, V]
}

func NewCache[K comparable, V any](size uint) *Cache[K, V] {
	return &Cache[K, V]{ //nolint:exhaustruct
		size:  size,
		inner: make(map[K]*el[K, V], size),
		lru:   newList[K, V](),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()

	if el, ok := c.inner[key]; ok {
		c.lru.toTop(el)
		el.value = value
		c.inner[key] = el
		c.mu.Unlock()

		return
	}

	if c.lru.size == c.size {
		el := c.lru.last()
		delete(c.inner, el.key)
		c.lru.remove(el)
	}

	el := c.lru.prepend(key, value)
	c.inner[key] = el
	c.mu.Unlock()
}

func (c *Cache[K, V]) Get(key K) (V, bool) { //nolint:ireturn
	c.mu.RLock()

	if el, ok := c.inner[key]; ok {
		c.lru.toTop(el)
		c.mu.RUnlock()

		return el.value, true
	}

	c.mu.RUnlock()

	var v V

	return v, false
}

func (c *Cache[K, V]) Del(key K) {
	c.mu.RLock()

	el, ok := c.inner[key]
	if !ok {
		c.mu.RUnlock()

		return
	}

	c.lru.remove(el)
	delete(c.inner, key)

	c.mu.RUnlock()
}
