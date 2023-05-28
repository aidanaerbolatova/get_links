package service

import (
	"fmt"
	"sync"
)

type Cache interface {
	Add(key, value string)
	Get(key string) (string, bool)
	Len() int
}

type CacheService struct {
	Archive sync.Map
}

func NewCacheService() *CacheService {
	return &CacheService{}
}

func (c *CacheService) Add(key, value string) {
	c.Archive.LoadOrStore(key, value)
}

func (c *CacheService) Get(key string) (string, bool) {
	temp, ok := c.Archive.Load(key)
	value := fmt.Sprintf("value: %v\n", temp)
	return value, ok
}

func (c *CacheService) Len() int {
	var i int
	c.Archive.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return i
}
