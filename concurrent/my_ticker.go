package main

import (
	"fmt"
	"time"
)

func main() {

	// 定时器执行多次
	ticker := time.NewTicker(time.Second * 2)

	for {
		<-ticker.C
		fmt.Println("时间到")
	}

}
