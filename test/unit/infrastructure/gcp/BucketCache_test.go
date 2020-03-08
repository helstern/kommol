package gcp

import (
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func runGet(cache gcp.BucketCache, nrTimes int32, postGetCount *int32, preGetCount *int32) {
	for i := int32(0); i < nrTimes; i++ {
		go func() {
			atomic.AddInt32(preGetCount, 1)
			cache.Get("a key")
			atomic.AddInt32(postGetCount, 1)
		}()
	}
}

func TestBucketCachedGetDoesNotBlocksOnRead(t *testing.T) {
	lock := sync.RWMutex{}
	cache := gcp.NewBucketCacheWithParams(make(map[string]gcp.BucketCacheItem), &lock)

	// use a high number of routines to make sure they end up on different threads
	var concurrentRoutines int32 = int32(3 * math.Pow10(3))
	var uPreGetCounter int32
	var uPostGetCounter int32
	lock.RLock()
	runGet(cache, concurrentRoutines, &uPostGetCounter, &uPreGetCounter)
	time.Sleep(1 * time.Second)
	if atomic.LoadInt32(&uPostGetCounter) != concurrentRoutines {
		t.Logf("post get counter := %d", uPostGetCounter)
		t.Errorf("should not block when reading")
	}
	lock.RUnlock()

}

func TestBucketCacheGetBlocksOnRead(t *testing.T) {
	lock := sync.RWMutex{}
	cache := gcp.NewBucketCacheWithParams(make(map[string]gcp.BucketCacheItem), &lock)

	// use a high number of routines to make sure they end up on different threads
	var concurrentRoutines int32 = int32(3 * math.Pow10(3))
	var lPreGetCounter int32
	var lPostGetCounter int32
	lock.Lock()
	runGet(cache, concurrentRoutines, &lPostGetCounter, &lPreGetCounter)

	// wait for completion
	time.Sleep(1 * time.Second)
	if atomic.LoadInt32(&lPreGetCounter) != concurrentRoutines {
		t.Errorf("should not block before reading")
	}
	if atomic.LoadInt32(&lPostGetCounter) != 0 {
		t.Logf("post get counter := %d", lPostGetCounter)
		t.Errorf("should block when writing")
	}
	lock.Unlock()

	// wait again for completion
	time.Sleep(1 * time.Second)
	if atomic.LoadInt32(&lPostGetCounter) != concurrentRoutines {
		t.Logf("post get counter := %d", lPostGetCounter)
		t.Errorf("should not block when reading")
	}
}
