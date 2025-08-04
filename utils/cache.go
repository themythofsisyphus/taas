package utils

import (
	"log"
	"strconv"
	"taas/config"

	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	Client *memcache.Client
}

func NewCacheClient(config *config.MemcacheConfig) *Cache {
	return &Cache{
		Client: memcache.New(config.Host + ":" + config.Port),
	}
}

func (c *Cache) Set(key string, value string) error {
	err := c.Client.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		log.Printf("[Memcache] set error: %v", err)
		return err
	}
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	item, err := c.Client.Get(key)

	if err != nil {
		log.Printf("[Memcache] get error: %v", err)
		return "", err
	}

	return string(item.Value), nil
}

func (c *Cache) Remove(key string) error {
	err := c.Client.Delete(key)
	if err != nil {
		log.Printf("[Memcache] delete error: %v", err)
		return err
	}
	return nil
}

func EntityCacheKey(name string, tenantID uint) string {
	return "ENTITY_KEY" + ":" + strconv.Itoa(int(tenantID)) + ":" + name
}
