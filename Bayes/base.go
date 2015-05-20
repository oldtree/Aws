package Bayes

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Flusher interface {
	Add(key string, value interface{})
	Remove(key string)
}

type dirtyElement struct {
	modified bool
	removed  bool
	key      string
	value    interface{}
}

type cacheItem struct {
	key   string
	value interface{}
}

type Cache struct {
	mu          sync.Mutex
	flushPeriod time.Duration
	data        map[string]*list.Element
	dirtyList   *list.List
	list        *list.List
	capacity    int
	flusher     Flusher
	maxNrDirty  int
}

func (c *Cache) Len() int {
	return len(c.data)
}

func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.flusher == nil {
		return
	}
	for e := c.dirtyList.Front(); e != nil; e = e.Next() {
		if de, ok := e.Value.(*dirtyElement); ok {
			if de.removed {
				c.flusher.Remove(de.key)
			} else if de.modified {
				c.flusher.Add(de.key, de.value)
			}
		}
	}
	c.dirtyList = list.New()
}

func (c *Cache) debug() {
	fmt.Printf("nr elems %v <= %v, nr dirty elems %v < %v\n",
		c.list.Len(), c.capacity, c.dirtyList.Len(), c.maxNrDirty)
	fmt.Println("-----------------elements------------")
	for e := c.list.Front(); e != nil; e = e.Next() {
		item := e.Value.(*cacheItem)
		fmt.Printf("%v: %v\n", item.key, item.value)
	}
	fmt.Println("-----------dirty elements------------")
	for e := c.dirtyList.Front(); e != nil; e = e.Next() {
		de := e.Value.(*dirtyElement)
		fmt.Printf("%v: %v; modified: %v; removed %v\n",
			de.key, de.value, de.modified, de.removed)
	}
	fmt.Println("-------------------------------------")
}

func (c *Cache) checkAndFlush() {
	c.mu.Lock()
	if c.maxNrDirty >= 0 && c.dirtyList.Len() >= c.maxNrDirty {
		c.mu.Unlock()
		c.Flush()
	} else {
		c.mu.Unlock()
	}
}

// capacity: nr elements in the cache.
// capacity < 0 means always in memory;
// capacity = 0 means no cache.
//
// maxNrDirty: < 0 means no flush.
//
// flushPeriod:
// flushPeriod > 1 second means periodically flush;
// flushPeriod = 0 second means no periodically flush;
// undefined in range (0, 1).
func New(capacity int, maxNrDirty int, flushPeriod time.Duration, flusher Flusher) *Cache {
	cache := new(Cache)

	cache.flushPeriod = flushPeriod
	cache.capacity = capacity
	if cache.capacity < 0 {
		cache.data = make(map[string]*list.Element, 1024)
	} else {
		cache.data = make(map[string]*list.Element, cache.capacity)
	}
	cache.list = list.New()
	cache.dirtyList = list.New()
	cache.flusher = flusher
	cache.maxNrDirty = maxNrDirty

	if flushPeriod.Seconds() > 0.9 {
		go func() {
			for {
				time.Sleep(flushPeriod)
				cache.Flush()
			}
		}()
	}
	return cache
}

func (c *Cache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.data[key]; ok {
		item := elem.Value.(*cacheItem)
		c.list.MoveToFront(elem)
		return item.value
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.checkAndFlush()
	defer c.mu.Unlock()

	de := &dirtyElement{
		modified: true,
		removed:  false,
		key:      key,
		value:    value,
	}

	if e, ok := c.data[key]; ok {
		item := e.Value.(*cacheItem)
		item.value = value
		c.list.MoveToFront(e)

		c.dirtyList.PushBack(de)
	} else {
		elem := c.list.PushFront(&cacheItem{key: key, value: value})
		c.data[key] = elem
		c.dirtyList.PushBack(de)
		if c.capacity >= 0 && len(c.data) > c.capacity {
			last := c.list.Back()
			item := last.Value.(*cacheItem)
			c.list.Remove(last)
			delete(c.data, item.key)
		}
	}
	return
}

func (c *Cache) Delete(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	de := &dirtyElement{
		modified: false,
		removed:  true,
		key:      key,
		value:    nil,
	}
	c.dirtyList.PushBack(de)
	if elem, ok := c.data[key]; ok {
		delete(c.data, key)
		item := elem.Value.(*cacheItem)
		return item.value
	}
	return nil
}