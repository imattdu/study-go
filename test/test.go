package test

import "fmt"

type Read struct {

}

func (r Read) Read(){
	fmt.Println("this is test Read read()")
}

func (r Read) Write(str string) string {
	return str
}
