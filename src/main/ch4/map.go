package main

import "fmt"

func main() {
	m1 := map[string]int{}
	m2 := make(map[string]int)

	m1["hello"] = 1
	m2["hello"] = 2

	alter1(m1, "hello", 10)
	alter2(m2, "hello", 20)

	fmt.Println(m1)
	fmt.Println(m2)

}

func alter1(m map[string]int, tag string, num int) {
	m[tag] = num
}

func alter2(m map[string]int, tag string, num int) {
	m[tag] = num
}
