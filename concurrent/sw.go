package main

import (
	"fmt"
	"sync"
)

var (
	wg1 sync.WaitGroup
)

func t1() {
	fmt.Println("t1...")
	// 计数器-2
	defer wg1.Done()
}

func main() {
	// 计数器+2
	wg1.Add(2)
	go add()
	go add()
	// 让main协程等待wg计数器为0在执行
	wg1.Wait()
}
