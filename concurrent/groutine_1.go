package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("is gr")
}

func main() {
	go hello()
	fmt.Println("main gr")
	time.Sleep(time.Second)
}
