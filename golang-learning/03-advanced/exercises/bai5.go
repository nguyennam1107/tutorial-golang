package main
import (
	"fmt"
	"sync"
	"time"
)
// ============= BÀI 5: GENERIC CACHE =============

// TODO: Generic Cache
type CacheItem[V any] struct {
	value      V
	expireTime time.Time
}

type Cache[K comparable, V any] struct {
	// Implement here
	items map[K]CacheItem[V]
	mu    sync.RWMutex
	maxSize int
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	// Implement here
	return &Cache[K, V]{
		items:  make(map[K]CacheItem[V]),
		maxSize: 0,
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	// Implement here
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem[V]{value: value}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	// Implement here
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.items[key]
	if !ok || item.expireTime.Before(time.Now()) {
		var zero V
		return zero, false
	}
	return item.value, true
}

func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration) {
	// Implement here
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem[V]{value: value, expireTime: time.Now().Add(ttl)}
}

func (c *Cache[K, V]) Delete(key K) {
	// Implement here
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache[K, V]) Clear() {
	// Implement here
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[K]CacheItem[V])
}

func (c *Cache[K, V]) Size() int {
	// Implement here
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

func (c *Cache[K, V]) SetMaxSize(maxSize int) {
	// Implement here
	c.maxSize = maxSize
}

func exercise5() {
	fmt.Println("\n=== BÀI 5: GENERIC CACHE ===")

	// TODO: Test basic operations
	cache := NewCache[string, int]()
	cache.Set("one", 1)
	cache.Set("two", 2)
	val, ok := cache.Get("one")
	fmt.Printf("Get 'one': %d, %v\n", val, ok)

	// TODO: Test TTL
	cache.SetWithTTL("temp", 42, 100*time.Millisecond)
	val, ok = cache.Get("temp")
	fmt.Printf("Before expiry: %d, %v\n", val, ok)
	time.Sleep(150 * time.Millisecond)
	val, ok = cache.Get("temp")
	fmt.Printf("After expiry: %d, %v\n", val, ok)
}

func main () {
	exercise5()
}