package main

import (
	"fmt"
)

func main() {

	intchan := make(chan int, 1000)
	primechan := make(chan int, 2000)
	exitchan := make(chan bool, 4)

	go putNum(intchan, 8000)
	for i := 0; i < 4; i++ {
		go primeNum(intchan, primechan, exitchan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitchan
		}
		close(exitchan)
		close(primechan)
	}()
	fmt.Println(len(primechan))
	for v := range primechan {
		fmt.Println(v)
	}

}

func putNum(intchan chan int, n int) {

	for i := 1; i <= n; i++ {
		intchan <- i
	}
	close(intchan)
}

func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {

	for {
		var num int
		var isSuccess bool
		var isSuShu bool
		num, isSuccess = <-intchan
		if isSuccess {
			isSuShu = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					isSuShu = false
				}
			}
			if isSuShu {
				primechan <- num
			}
		} else {
			exitchan <- true
			break
		}
	}
}
