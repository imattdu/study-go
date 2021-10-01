package main

import (
	"fmt"
	"stugo/test"
)

type Read interface {
	Read()
}

type Write interface {
	Write(str string) string
}

// 接口组合
type IO interface {
	Read
	Write
}

func main()  {

	var read Read
	read = test.Read{}
	//// 大写才可以被访问
	//read.Read()
	//
	//fmt.Printf("%T %v\n", read, read)
	//
	//switch v := read.(type) {
	//case test.Read:
	//	fmt.Println(v)
	//}

	// type assertion
	// 类型断言
	fmt.Println("ok")
	// inter face 一个实现者类型 一个实现者值
	r1, OK  := read.(test.Read)
	if OK {
		r1.Read()
	} else {
		fmt.Println("出错了")
	}

	var t1 IO = test.Read{}
	fmt.Println(t1.Write("HELLO WORD"))

}
