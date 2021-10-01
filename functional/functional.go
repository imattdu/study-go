package main

import "fmt"

// 函数闭包 返回一个函数
func adder() func(int) int {
	var sum int = 0
	return func(v int) int {
		sum += v
		return sum
	}
}


func main(){

	add := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}

}


