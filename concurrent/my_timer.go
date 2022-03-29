package main

import (
	"fmt"
	"time"
)

func main() {
	// 只执行一次
	fmt.Println(time.Now(), "begin")
	timer := time.NewTimer(time.Second * 2)
	fmt.Println(<-timer.C, "时间到")
	fmt.Println(time.Now(), "end....")
}
