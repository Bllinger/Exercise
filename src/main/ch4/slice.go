package main

import (
	"fmt"
)

func main() {
	var num = [9]int{1, 2, 34, 5, 6, 7, 8, 9, 10}

	s1 := num[4:7]
	fmt.Println(s1)
	_ = append(s1, 100)
	//nan := math.NaN()

	fmt.Println(num)

	fmt.Println(s1[:5])

	var s2 []int
	fmt.Println(s2 == nil)
	s3 := []int{}
	fmt.Println(s3 == nil)
}
