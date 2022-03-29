package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := make(chan int, 10)
	//t2 := make(chan int, 10)

	// 使用协程 range 就不需要close
	go func() {
		for t := range t1 {
			fmt.Println(t)
		}
	}()

	for i := 0; i < 10; i++ {
		t1 <- i
	}
	time.Sleep(time.Second * 2)

}
