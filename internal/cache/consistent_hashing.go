package cache

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func(data []byte) uint32

type ConsistentHash struct {
	hash     HashFunc
	replicas int
	keys     []int
	hashMap  map[int]string
}

func NewConsistentHash(replicas int, fn HashFunc) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		hash:     fn,
		keys:     make([]int, 0),
		hashMap:  make(map[int]string),
	}
}

func (c *ConsistentHash) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < c.replicas; i++ {
			hash := int(c.hash([]byte(strconv.Itoa(i) + node)))
			c.keys = append(c.keys, hash)
			c.hashMap[hash] = node
		}
	}
	sort.Ints(c.keys)
}

func (c *ConsistentHash) Get(key string) string {
	if len(c.keys) == 0 {
		return ""
	}
	hash := int(c.hash([]byte(key)))
	idx := sort.Search(len(c.keys), func(i int) bool { return c.keys[i] >= hash })
	if idx == len(c.keys) {
		idx = 0
	}
	return c.hashMap[c.keys[idx]]
}

func Hash(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
