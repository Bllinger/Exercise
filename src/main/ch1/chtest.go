package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	for i := 1; i < 4; i++ {
		go add(i, ch)
	}

	for i := 1; i < 4; i++ {
		fmt.Println(<-ch)
	}
}

func add(num int, ch chan int) {
	//if num == 1 {
	//	time.Sleep(time.Duration(1)*time.Second)
	//}

	ch <- num
}
