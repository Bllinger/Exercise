package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma1("12345"))

	fmt.Printf("%08b\n", 1&^2)
}

func comma1(s string) string {
	var buf bytes.Buffer
	l := len(s)
	buf.WriteString("," + s[l-3:])

	for l -= 3; (l - 3) > 0; l -= 3 {
		temp := buf.String()
		buf.Reset()
		buf.WriteString("," + s[l-3:l] + temp)
	}

	return s[0:l] + buf.String()
}
