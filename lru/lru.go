package lru

import (
	"sync"

	"github.com/dolthub/swiss"
)

type Cache[K comparable, V any] struct {
	size  uint
	mu    sync.RWMutex
	inner *swiss.Map[K, *el[K, V]]
	lru   *list[K, V]
}

func NewCache[K comparable, V any](size uint) *Cache[K, V] {
	return &Cache[K, V]{
		size:  size,
		inner: swiss.NewMap[K, *el[K, V]](uint32(size)),
		lru:   newList[K, V](),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()

	if el, ok := c.inner.Get(key); ok {
		c.lru.toTop(el)
		el.value = value
		c.inner.Put(key, el)
		c.mu.Unlock()

		return
	}

	if c.lru.size == c.size {
		el := c.lru.last()
		c.inner.Delete(el.key)
		c.lru.remove(el)
	}

	el := c.lru.prepend(key, value)
	c.inner.Put(key, el)
	c.mu.Unlock()
}

func (c *Cache[K, V]) Get(key K) (V, bool) { //nolint:ireturn
	c.mu.RLock()

	if el, ok := c.inner.Get(key); ok {
		c.lru.toTop(el)
		c.mu.RUnlock()

		return el.value, true
	}

	c.mu.RUnlock()

	var v V

	return v, false
}
