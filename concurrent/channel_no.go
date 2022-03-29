package main

import "fmt"

func main() {
	//c1 := make(chan int)
	//c1 <- 1
	//<-c1

	c2 := make(chan int)
	go func() {
		c2 <- 1
	}()
	fmt.Println(<-c2)

	c3 := make(chan int, 10)
	c3 <- 11
	fmt.Println(len(c3), cap(c3))
}
