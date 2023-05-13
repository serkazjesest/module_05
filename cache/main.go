package main

import "time"

type Cache struct {
	storage map[string]int
}

func (c *Cache) Increase(key string, value int) {
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	delete(c.storage, key)
}
