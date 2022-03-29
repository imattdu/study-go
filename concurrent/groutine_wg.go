package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func w1() {
	defer wg.Done()
	fmt.Println("wg1")
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go w1()
	}
	wg.Wait()
	fmt.Println("main")
}
