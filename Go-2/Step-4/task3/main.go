package main

import "sync"

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

func (c *ConcurrentQueue) Enqueue(element interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = append(c.queue, element)
}

func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	element := c.queue[0]
	c.queue = c.queue[1:]
	return element
}
