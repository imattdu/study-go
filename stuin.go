package main

import (
	"fmt"
	"stugo/stuinterface"
)

func main()  {

	var i1 stuinterface.I = stuinterface.I{}
	i1.Name = "hello"
	fmt.Println(i1.String())

}
