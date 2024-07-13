package cache

type Cache interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
}

type CacheWithHash struct {
	cache          Cache
	consistentHash *ConsistentHash
}

func NewCacheWithHash(cache Cache, consistentHash *ConsistentHash) *CacheWithHash {
	return &CacheWithHash{
		cache:          cache,
		consistentHash: consistentHash,
	}
}

func (c *CacheWithHash) Get(key string) (interface{}, bool) {
	node := c.consistentHash.Get(key)
	return c.cache.Get(node + key)
}

func (c *CacheWithHash) Put(key string, value interface{}) {
	node := c.consistentHash.Get(key)
	c.cache.Put(node+key, value)
}
