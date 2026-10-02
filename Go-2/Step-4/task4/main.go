package main

import "sync"

var (
	Buf   []int
	mutex sync.Mutex
)

func Write(num int) {
	mutex.Lock()
	defer mutex.Unlock()
	Buf = append(Buf, num)
}

func Consume() int {
	mutex.Lock()
	defer mutex.Unlock()
	res := Buf[0]
	Buf = Buf[1:]
	return res
}
