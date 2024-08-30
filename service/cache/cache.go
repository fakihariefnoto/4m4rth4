package cache

import (
	"errors"
	"time"
)

type (
	Cache struct {
		data   map[string]data
		ttl    int
		init   bool
		create bool
	}

	data struct {
		d interface{}
		t time.Time
	}
)

// Init expire cache in second
func InMemory(expire int) *Cache {
	return &Cache{
		data:   make(map[string]data),
		ttl:    expire,
		create: true,
	}
}

// Cap capacity
func (c *Cache) Cap() int {
	return len(c.data)
}

// InitGC to initiate GC cycle, duration should be in milisecond
func (c *Cache) InitGC(d int) {
	c.init = true
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(d))
			go func() {
				c.startGC()
			}()
		}
	}()
}

// startGC will collect all expired key and delete it
func (c *Cache) startGC() {
	now := time.Now()
	for key, value := range c.data {
		if now.After(value.t.Add(time.Second * time.Duration(c.ttl))) {
			delete(c.data, key)
		}
	}
}

// Set to add new Cache or update existing with same type data
func (c *Cache) Set(key string, value interface{}) error {
	if !c.init && !c.create {
		return errors.New("error init first")
	}
	c.data[key] = data{
		d: value,
		t: time.Now(),
	}
	return nil
}

// Get to retrieve data based on key
func (c *Cache) Get(key string) (interface{}, error) {
	if !c.init && !c.create {
		return nil, errors.New("error init first")
	}

	if data, val := c.data[key]; val {
		now := time.Now()
		if now.After(data.t.Add(time.Second * time.Duration(c.ttl))) {
			return nil, errors.New("expired")
		}
		return data.d, nil
	}

	return nil, errors.New("not found")
}
