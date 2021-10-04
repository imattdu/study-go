package main

import (
	"fmt"
	"time"
)

// 有自己的阻塞机制
func main() {

	intchan := make(chan int, 10)
	exitchan := make(chan bool, 1)
	go writeData(intchan)
	go readData(intchan, exitchan)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}
}

func writeData(intchan chan int) {
	for i := 0; i < 50; i++ {
		intchan <- i
		fmt.Println("写入数据", i)
		time.Sleep(200)
	}
	close(intchan)
}

func readData(intchan chan int, exitchan chan bool) {
	for {
		/*if len(exitchan) != 0 {
			close(intchan)
			for v := range intchan {
				fmt.Println(v)
			}
			close(exitchan)
			break
		}*/
		v, ok := <-intchan
		if !ok {
			break
		} else {
			fmt.Println("读取数据", v)
			time.Sleep(time.Second * 2)
		}
	}
	exitchan <- true
	close(exitchan)
}
