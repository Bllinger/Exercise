package main

import "fmt"

func main() {
	s := "a/b/c"
	flag := false
	var right, left = len(s), 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' && !flag {
			right = i
			flag = true
		}

		if flag && s[i] == '/' {
			left = i + 1
			break
		}
	}

	if !flag {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '/' {
				left = i + 1
				break
			}
		}
	}

	fmt.Println(s[left:right])
}
