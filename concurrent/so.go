package main

import (
	"fmt"
	"sync"
)

var (
	so sync.Once
	m  map[int]string
)

func a2() {
	m = make(map[int]string)
}

func a1() {
	// 不使用 so 通过判断 m != nil 可能没有初始化完成但是 != nil
	// 使用协程又会有安全问题
	so.Do(a2)
	fmt.Println("this is good")
}

func main() {

}
