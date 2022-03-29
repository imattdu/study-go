package main

func main() {
	t1 := make(chan<- int, 5)
	t1 <- 1
	// 报错
	//fmt.Println(<-t1)

}
