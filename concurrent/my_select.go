package main

import (
	"fmt"
	"math/rand"
)

func main() {

	t1 := make(chan int, 10)
	t2 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		t1 <- i
		t2 <- rand.Int()
	}
b1:
	for {
		select {
		case v := <-t1:
			fmt.Println(v)
		case v := <-t2:
			fmt.Println("t2:", v)
		default:
			fmt.Println("else")
			break b1
		}
	}

}
