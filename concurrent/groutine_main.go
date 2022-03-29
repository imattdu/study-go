package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("start...")
		time.Sleep(time.Second * 10)
		fmt.Println("end...")
	}()

	fmt.Println("main start...")
	time.Sleep(time.Second * 5)
	fmt.Println("main end...")
}
