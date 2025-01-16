package discovery

import (
	"testing"
)

// TestCacheSet tests the Set function
func TestCacheGetSet(t *testing.T) {

	cache := NewCache(10)

	cache.Set("key1", "value1")

	cacheItem := cacheItem{
		key:   "key1",
		value: "value1",
	}

	if cache.list.Len() != 1 {
		t.Errorf("Expected cache to have 1 item, got %d", cache.list.Len())
	}

	if cache.Get("key1") != cacheItem {
		t.Errorf("Expected value to be 'value1', got %s", cache.Get("key1"))
	}

}

// TestCacheEvict tests the evict function

func TestCacheEvict(t *testing.T) {

	cache := NewCache(2)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	if cache.list.Len() != 2 {
		t.Errorf("Expected cache to have 2 items, got %d", cache.list.Len())
	}

	if cache.Get("key1") != nil {
		t.Errorf("Expected key1 to be evicted, got %s", cache.Get("key1"))
	}

}
