package cache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[string]*list.Element
	mutex    sync.Mutex
}

type entry struct {
	key   string
	value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (lru *LRUCache) Get(key string) (interface{}, bool) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

func (lru *LRUCache) Put(key string, value interface{}) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
	} else {
		if len(lru.cache) >= lru.capacity {
			lru.evict()
		}
		newElem := lru.list.PushFront(&entry{key, value})
		lru.cache[key] = newElem
	}
}

func (lru *LRUCache) evict() {
	if lru.list.Len() > 0 {
		oldest := lru.list.Back()
		if oldest != nil {
			delete(lru.cache, oldest.Value.(*entry).key)
			lru.list.Remove(oldest)
		}
	}
}
