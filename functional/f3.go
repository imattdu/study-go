package main

import "fmt"

func t1() {
	t2(func() {
		fmt.Println("t1 s")
	})
}

func t2(f func()) {
	fmt.Println("t2")
	f()
}

func main()  {
	t1()
	t2(func() {
		fmt.Println("hello word")
	})
}
