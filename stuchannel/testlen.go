package main

import "fmt"

func main() {
	var intchan chan int = make(chan int, 10)
	intchan <- 1
	intchan <- 2
	close(intchan)
	fmt.Println(len(intchan))
}
