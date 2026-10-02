package main

import "sync"

type Counter struct {
	value int
	mu    sync.RWMutex
}

type Count interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения

}

func (c *Counter) Increment() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
