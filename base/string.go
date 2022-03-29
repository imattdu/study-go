package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	if i, err := strconv.Atoi("12343"); err == nil {
		fmt.Println(i, unsafe.Sizeof(i))
	}

}
