package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

type Meter float64
type Inch float64
type Pound float64
type Kg float64

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
func (i Inch) String() string {
	return fmt.Sprintf("%gin", i)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (kg Kg) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func main() {
	var choose int
	var num float64

	fmt.Println("请选择换算的单位：(输入数字)")
	fmt.Println("1.温度")
	fmt.Println("2.长度")
	fmt.Println("3.重量")
	fmt.Scanln(&choose)

	if len(os.Args[1:]) < 1 {
		flag := 1

		for flag == 1 {
			fmt.Print("请输入待换算待数字:")
			fmt.Scanln(&num)

			chooseF(choose, num)

			fmt.Print("是否需要继续换算，是（1）否（0）")
			fmt.Scanln(&flag)
		}
	} else {
		for _, p := range os.Args[1:] {
			num, err := strconv.ParseFloat(p, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}

			chooseF(choose, num)
		}
	}
}

func chooseF(c int, n float64) {
	switch c {
	case 1:
		toTemperature(n)
		toTemperature(n)
	case 2:
		toLength(n)
	case 3:
		toHeight(n)
	default:
		fmt.Println("输入数字不合法")
	}
}
func toTemperature(n float64) {
	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)

	fmt.Printf("%s = %s\t%s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func toLength(n float64) {
	m := Meter(n)
	i := Inch(n)

	fmt.Printf("%s = %s\t%s = %s\n", m, MToI(m), i, IToM(i))
}

func MToI(m Meter) Inch {
	return Inch(m * 100 * 0.3937)
}

func IToM(i Inch) Meter {
	return Meter(i * 2.54 * 0.01)
}

func toHeight(n float64) {
	p := Pound(n)
	k := Kg(n)

	fmt.Printf("%s = %s\t%s = %s\n", p, PToK(p), k, KToP(k))
}

func PToK(p Pound) Kg {
	return Kg(p * 0.45359237)
}

func KToP(k Kg) Pound {
	return Pound(k * 2.2046226218488)
}
