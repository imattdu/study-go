package main

import "fmt"

func main() {
	c1 := make(chan int)
	select {
	case c1 <- 1:
		fmt.Println("matt")
	default:
		fmt.Println("no")
	}
}
