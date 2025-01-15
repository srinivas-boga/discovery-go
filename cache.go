package discovery

import (
	"container/list"
	"sync"
)

// LRU cache
type Cache struct {
	instances sync.Map
	capacity  int
	list      *list.List
}

type cacheItem struct {
	key   string
	value interface{}
}

// NewCache creates a new Cache
func NewCache(capacity int) *Cache {
	return &Cache{
		instances: sync.Map{},
		capacity:  capacity,
		list:      list.New(),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) interface{} {
	if value, ok := c.instances.Load(key); ok {
		c.list.MoveToFront(value.(*list.Element))
		return value.(*list.Element).Value
	}
	return nil
}

// Set adds a value to the cache
func (c *Cache) Set(key string, value interface{}) {

	if c.list.Len() >= c.capacity {
		c.evict()
	}

	newCacheItem := cacheItem{
		key:   key,
		value: value,
	}

	element := c.list.PushFront(newCacheItem)
	c.instances.Store(key, element)

}

// evict removes the least recently used element from the cache
func (c *Cache) evict() {

	element := c.list.Back()

	c.list.Remove(element)
	c.instances.Delete(element.Value.(cacheItem).key)

}
