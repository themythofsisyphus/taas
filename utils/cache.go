// Package utils contains utility functions including caching logic and helpers.
package utils

import (
	"log"
	"strconv"
	"taas/config"

	"github.com/bradfitz/gomemcache/memcache"
)

// Cache wraps the Memcache client.
type Cache struct {
	Client *memcache.Client
}

// NewCacheClient creates a new memcache client using the provided config.
func NewCacheClient(cfg *config.MemcacheConfig) *Cache {
	return &Cache{
		Client: memcache.New(cfg.Host + ":" + cfg.Port),
	}
}

// Set adds a key-value pair to the memcache.
func (c *Cache) Set(key, value string) error {
	err := c.Client.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		log.Printf("[Memcache][Set] Error setting key %s: %v", key, err)
		return err
	}
	return nil
}

// Get retrieves a value from memcache using the given key.
func (c *Cache) Get(key string) (string, error) {
	item, err := c.Client.Get(key)
	if err != nil {
		log.Printf("[Memcache][Get] Error retrieving key %s: %v", key, err)
		return "", err
	}
	return string(item.Value), nil
}

// Remove deletes a key from memcache.
func (c *Cache) Remove(key string) error {
	err := c.Client.Delete(key)
	if err != nil {
		log.Printf("[Memcache][Remove] Error deleting key %s: %v", key, err)
		return err
	}
	return nil
}

// EntityCacheKey generates a namespaced cache key for entity data per tenant.
func EntityCacheKey(name string, tenantID uint) string {
	return "ENTITY_KEY:" + strconv.Itoa(int(tenantID)) + ":" + name
}
