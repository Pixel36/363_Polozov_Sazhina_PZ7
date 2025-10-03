package main

import (
 "fmt"
 "time"
)

type cacheItem struct {
 value      interface{}
 expiration int64
}

type Cache struct {
 items map[string]cacheItem
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
 if c.items == nil {
  c.items = make(map[string]cacheItem)
 }
 c.items[key] = cacheItem{value, time.Now().Add(ttl).Unix()}
}
func (c *Cache) Get(key string) (interface{}, bool) {
 item, ok := c.items[key]
 if !ok || time.Now().Unix() > item.expiration {
  return nil, false
 }
 return item.value, true
}

func (c *Cache) Delete(key string) {
 delete(c.items, key)
}

func main() {
 cache := Cache{}
 cache.Set("name", "Go", 2*time.Second)
 val, ok := cache.Get("name")
 fmt.Println(val, ok)
 time.Sleep(3 * time.Second)
 val, ok = cache.Get("name")
 fmt.Println(val, ok)
}
