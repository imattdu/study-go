package main

import (
	"fmt"
	"sync"
)

var (
	cnt    int
	lock   sync.Mutex
	rwLock sync.RWMutex
	w      sync.WaitGroup
)

func add() {

	for i := 0; i < 1000; i++ {
		lock.Lock()
		cnt++
		lock.Unlock()
	}
	w.Done()
}

func get() {
	// 读锁
	rwLock.RLock()
	rwLock.RUnlock()
	// 写锁
	rwLock.Lock()
	rwLock.Unlock()
}

func main() {
	// 加锁 2000 不加锁 1220
	w.Add(2)
	go add()
	go add()
	w.Wait()
	fmt.Println(cnt)

}
