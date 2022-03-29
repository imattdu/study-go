package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int64 = 100
	res := atomic.AddInt64(&i, 100)
	res = atomic.SwapInt64(&i, 3)
	fmt.Println(res, i)

}
