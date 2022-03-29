package main

import (
	"fmt"
)

func main() {

	//var t1 chan int = make(chan int, 5)
	//// var t2 chan int = make(chan int)
	//t1 <- 11
	//fmt.Println(<- t1)
	//close(t1)
	//t2 := make(chan int)
	//go func() {
	//	fmt.Println(len(t2))
	//	fmt.Println(cap(t2))
	//	<- t2
	//
	//}()
	//
	//t2 <- 1

	t3 := make(chan int, 100)
	for i := 0; i < 1; i++ {
		t3 <- i
	}
	// close(t3)
	<-t3
	_, ok := <-t3
	fmt.Println(ok)

	v, ok := <-t3
	fmt.Println(ok, v)
	close(t3)
	// range 需要先关闭
	for v := range t3 {
		fmt.Println(v)
	}

	// 通道没有关闭 ok一直为true所以报错
	for {
		v, ok := <-t3
		fmt.Println(ok)
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

}
