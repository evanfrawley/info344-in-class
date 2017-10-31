package main

import (
	"time"
	"sync"
	"fmt"
)

type entry struct {
	value     string
	expiresAt time.Time
}

//Cache is a TTL cache that is safe for concurrent use
type Cache struct {
	entries map[string]*entry
	//TODO: protect this for concurrent use!
	mx sync.RWMutex
}

//NewCache constructs a new Cache object
func NewCache() *Cache {
	c :=  &Cache{
		entries: map[string]*entry{},
	}
	go c.janitor()
	return c
}

//Set adds a key/value to the cache
func (c *Cache) Set(key string, value string, timeToLive time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[key] = &entry{value, time.Now().Add(timeToLive)}
}

//Get gets the value associated with a key
func (c *Cache) Get(key string) (string, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	//TODO: implement this
	entry, isThere := c.entries[key]
	if !isThere {
		return "", false
	}
	return entry.value, true
}

func (c *Cache) janitor() {
	for {
		time.Sleep(time.Second)
		now := time.Now()
		c.mx.Lock()

		fmt.Println("janitor is running")
		for key, entry := range c.entries {
			if entry.expiresAt.Before(now) {
				fmt.Printf("purging key %s\n", key)
				delete(c.entries, key)
			}
		}
		c.mx.Unlock()
	}
}

