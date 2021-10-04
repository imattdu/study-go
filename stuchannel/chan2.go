package main

import "fmt"

type Cat struct {
	Name string
}

func main() {
	var allchan chan interface{} = make(chan interface{}, 10)

	var cat1 Cat = Cat{Name: "matt"}
	allchan <- cat1
	allchan <- 1
	cat2 := <-allchan
	realcat2 := cat2.(Cat)
	fmt.Println(realcat2.Name)
}
