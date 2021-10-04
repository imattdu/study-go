package main

import "fmt"

func main() {
	var intchan chan int = make(chan int, 10)
	intchan <- 1
	intchan <- 2
	intchan <- 3
	intchan <- 4
	intchan <- 5

	close(intchan)

	for v := range intchan {
		fmt.Println("intchan:", v)
	}

}
