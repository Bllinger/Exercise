package main

import (
	"fmt"
	"tempconv"
)

func main() {
	c := tempconv.Celsius(24)

	fmt.Println(c.String())
	fmt.Println("Celsius of Boiling:", tempconv.BoilingC)
	fmt.Println("Fahrenheit of AbsoluteZero:", tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println("Kelvins of Zero:", tempconv.KtoC(tempconv.ZeroK))
}
