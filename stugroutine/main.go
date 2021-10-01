package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("HELLO WORD")
		time.Sleep(time.Second)
	}
}

func main()  {

	runtime.GOMAXPROCS(8)
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("HELLO")
		time.Sleep(time.Second)
	}
}
