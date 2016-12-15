package main

import "time"

// Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and set.
//
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// set(key, value) - Set or insert the value if the key is not already present.

// When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item.
// For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently *used* key would be evicted.

type LFUCache struct {
	Capacity int
	Size     int
	Table    map[int]*Item // key, val
}

type Item struct {
	Value   int
	Freq    int
	Recency time.Time
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		Size:     0,
		Table:    make(map[int]*Item, capacity),
	}
}

func (cache *LFUCache) Get(key int) int {
	// Key exists.
	if item, ok := cache.Table[key]; ok {
		// Update the frequency and recency.
		cache.Table[key] = &Item{
			Value:   item.Value,
			Freq:    item.Freq + 1,
			Recency: time.Now(),
		}
		return item.Value
	}
	return -1
}

func (cache *LFUCache) Set(key int, value int) {
	if cache.Capacity == 0 {
		return
	}
	// Key does not exist.
	if item, ok := cache.Table[key]; !ok {
		// If we reach capacity:
		// find and evict least frequently & recently used element.
		if cache.Capacity == cache.Size {
			cache.findAndEvict()
		}

		// Add a new k,v pair.
		cache.Table[key] = &Item{
			Value:   value,
			Freq:    1,
			Recency: time.Now(),
		}

		// Cache size keeps increasing till it reaches the capacity.
		if cache.Size < cache.Capacity {
			cache.Size++
		}

	} else {
		// Key exists.
		cache.Table[key] = &Item{
			Value:   value,
			Freq:    item.Freq + 1,
			Recency: time.Now(),
		}

	}
}

func (cache *LFUCache) findAndEvict() {
	victim := &Item{
		Value:   -1,
		Freq:    1<<31 - 1, // MaxInt32
		Recency: time.Now(),
	}

	victimKey := -1
	for k, v := range cache.Table {
		if victim.Freq > v.Freq {
			victim = v
			victimKey = k
		} else if victim.Freq == v.Freq {
			if victim.Recency.After(v.Recency) {
				victim = v
				victimKey = k
			}
		}
	}

	// Remove victim element from table.
	if victimKey >= 0 {
		delete(cache.Table, victimKey)
	}
}
