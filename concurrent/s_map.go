package main

import (
	"fmt"
	"sync"
)

func main() {
	m1 := sync.Map{}

	m1.Store("name", "matt")
	m1.Store("age", 22)

	fmt.Println(m1.Load("name"))

	v, ok := m1.LoadOrStore("name", "matt")
	if ok {
		fmt.Println(v)
	}

	m1.Delete("name")

	// 存在返回false
	v, ok = m1.LoadOrStore("name", "aaa")
	if !ok {
		fmt.Println(v, "2222load and store")
	}

	//m1.Delete("age")

	// 不存在返回 false
	v, ok = m1.Load("ccc")
	fmt.Println(v, ok)

	m1.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
