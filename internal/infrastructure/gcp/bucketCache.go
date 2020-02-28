package gcp

import "cloud.google.com/go/storage"

type BucketCacheItem struct {
	website *storage.BucketWebsite
}

type BucketCache struct {
	items map[string]BucketCacheItem
}

func (this *BucketCache) Get(key string) (*storage.BucketWebsite, bool) {
	existingItem, exists := this.items[key]
	if exists {
		return existingItem.website, exists
	}

	return nil, exists
}

func (this *BucketCache) Put(key string, item *storage.BucketWebsite) *storage.BucketWebsite {
	existingItem, exists := this.Get(key)
	cacheItem := BucketCacheItem{
		website: item,
	}
	this.items[key] = cacheItem

	if exists {
		return existingItem
	}
	return nil
}
