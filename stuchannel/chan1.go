package main

import "fmt"

func main() {
	var intchan chan int = make(chan int, 3)

	fmt.Println(intchan)

	intchan <- 1
	num := 11
	intchan <- num
	intchan <- 985

	fmt.Println(len(intchan), cap(intchan))

	fmt.Println(<-intchan)
	fmt.Println(<-intchan)
	fmt.Println(len(intchan), "len cap")

}
