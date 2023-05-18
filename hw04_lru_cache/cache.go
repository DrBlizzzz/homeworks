package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	inCache := false
	for k := range cache.items {
		if k == key {
			inCache = true
		}
	}
	if inCache {
		cache.queue.MoveToFront(cache.items[key])
		node := cache.queue.Front()
		node.Value = value.(int)
		cache.items[key] = node
	} else {
		if cache.queue.Len() >= cache.capacity {
			search_val := cache.queue.Back().Value.(int)
			for k, v := range cache.items {
				if v.Value.(int) == search_val {
					delete(cache.items, k)
				}
			}
			cache.queue.Remove(cache.queue.Back())

		}
		cache.queue.PushFront(value.(int))
		cache.items[key] = cache.queue.Front()
	}

	return inCache
}

func (cache *lruCache) Get(key Key) (v interface{}, inCache bool) {
	inCache = false
	for k, val := range cache.items {
		if k == key {
			v = val.Value
			inCache = true
		}
	}
	if !inCache {
		return nil, false
	}
	cache.queue.MoveToFront(cache.items[key])
	cache.items[key] = cache.queue.Front()
	return v, inCache
}

func (cache *lruCache) Clear() {
	for cache.queue.Len() != 0 {
		to_search := cache.queue.Front().Value.(int)
		for k, v := range cache.items {
			if v.Value.(int) == to_search {
				delete(cache.items, k)
			}
		}
		cache.queue.Remove(cache.queue.Front())
	}
}
