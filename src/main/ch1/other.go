package main

import "fmt"

func main() {
	switch test() {
	case 1:
		fmt.Println("into 1!")
	case 2:
		fmt.Println("into 2")
	case 3:
		fmt.Println("into 3")
	default:
		fmt.Println("into default")

	}

	fmt.Println("hello mac")
}

func test() int {
	return 1
}
