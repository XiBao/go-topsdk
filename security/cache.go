package security

import (
	"sync"
)

type ICache interface {
	Get(k string) *Context
	Set(k string, v *Context)
}

type Cache struct {
	data map[string]*Context
	m    *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]*Context),
		m:    new(sync.RWMutex),
	}
}

func (c *Cache) Get(k string) *Context {
	c.m.RLock()
	v, ok := c.data[k]
	c.m.RUnlock()
	if !ok {
		return nil
	}

	return v
}

func (c *Cache) Set(k string, v *Context) {
	c.m.Lock()
	c.data[k] = v
	c.m.Unlock()
}
