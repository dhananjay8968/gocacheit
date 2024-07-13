package main

import (
    "gocacheit/internal/cache"
    "gocacheit/internal/server"
)

func main() {
    lruCache := cache.NewLRUCache(100) 
    hash := cache.NewConsistentHash(10, cache.Hash)
    cacheWithHash := cache.NewCacheWithHash(lruCache, hash)

    cacheServer := server.NewCacheServer(cacheWithHash)
    cacheServer.Start("8080") 
}
