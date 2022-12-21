package cacheprovider

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

var myCache CacheItf
var Cacheset *cache.Cache

type CacheItf interface {
	Set(key string, token string, expiration time.Duration) bool
	Get(key string) (string, bool)
	Delete(key string) bool
	Close()
}
type AppCache struct {
	client *cache.Cache
}

// SetProvider sets specific in memory provider e.g, redis, mem cache etc...
func SetProvider(provider CacheItf) error {
	if provider == nil {
		return errors.New("provider can not be nil")
	}
	if myCache != nil {
		myCache.Close()
	}
	myCache = provider
	return nil
}

// GetMemoryStore returns the memory store with default as Cache provider
func GetMemoryStore() CacheItf {
	if myCache != nil {
		return myCache
	}
	myCache = NewCacheProvider()
	return myCache
}
func NewCacheProvider() *AppCache {
	return &AppCache{
		client: cache.New(10*time.Minute, 10*time.Minute),
	}
}
func (r *AppCache) Set(key string, token string, expiration time.Duration) bool {

	r.client.Set(key, token, expiration)
	return true
}
func (r *AppCache) Get(key string) (string, bool) {
	var catch string
	var found bool
	data, found := r.client.Get(key)
	if found {
		catch = data.(string)
	}
	return catch, found
}
func (r *AppCache) Delete(key string) bool {
	r.client.Delete(key)
	return true
}
func (r *AppCache) Close() {
	r.Close()
}
func CacheGet() *cache.Cache {
	return Cacheset
}
