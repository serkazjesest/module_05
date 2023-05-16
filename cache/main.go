package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

const (
	k1   = "key1"
	step = 7
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	cache := Cache{storage: make(map[string]int)}
	semaphore := make(chan int, 4)

	for i := 0; i < 10; i++ {
		select {
		case semaphore <- i:
			go func() {
				cache.Increase(k1, step)
			}()
			cancel()
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			fmt.Println(cache.Get(k1))
			return
		default:
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 20)
		}
	}

	for i := 0; i < 10; i++ {
		select {
		case semaphore <- i:
			go func(i int) {
				cache.Set(k1, step*i)
			}(i)
			cancel()
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			fmt.Println(cache.Get(k1))
			return
		default:
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 20)
		}
	}
}
