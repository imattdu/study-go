package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1 map[int]int = make(map[int]int, 10)
	// 互斥锁
	lock sync.Mutex
)

func test(i int) {
	var res int = 1
	for t := 1; t <= i; t++ {
		res *= t
	}
	lock.Lock()
	m1[i] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(10 * time.Second)
	// 主线程并不知道啥时候结束
	// 目前不知道什么原因为啥会警告
	lock.Lock()
	for index, val := range m1 {
		fmt.Println(index, val)
	}
	lock.Unlock()
}
