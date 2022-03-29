package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c1 <- i
	}
	go func() {
		for v := range c1 {
			fmt.Println(v)
		}
	}()
	fmt.Println("hello word")
	time.Sleep(time.Second * 1)

}
