package gcp

import (
	"cloud.google.com/go/storage"
	"sync"
)

type BucketCacheItem struct {
	website *storage.BucketWebsite
}

type BucketCache struct {
	items map[string]BucketCacheItem
	lock  sync.RWMutex
}

func (this *BucketCache) Get(key string) (*storage.BucketWebsite, bool) {

	this.lock.RLock()
	existingItem, exists := this.items[key]
	this.lock.RUnlock()

	if exists {
		return existingItem.website, exists
	}
	return nil, exists
}

func (this *BucketCache) Put(key string, item *storage.BucketWebsite) {
	cacheItem := BucketCacheItem{
		website: item,
	}

	this.lock.Lock()
	this.items[key] = cacheItem
	this.lock.Unlock()
}

func NewBucketCache() BucketCache {
	return BucketCache{
		items: make(map[string]BucketCacheItem),
		lock:  sync.RWMutex{},
	}
}
