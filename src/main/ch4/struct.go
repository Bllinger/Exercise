package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p := new(Point)
	add1(p, 1, 2)
	fmt.Println(p)

	p1 := Point{}
	add2(p1, 2, 3)
	fmt.Println(p1)
	var p2 = Point{1, 2}
	fmt.Println(p2)
	pp1 := &Point{}
	fmt.Println(*pp1)
	pp2 := new(Point)
	*pp2 = Point{}
}

func add2(p Point, x, y int) {
	p.x = x
	p.y = y
}

func add1(p *Point, x, y int) {
	p.x = x
	p.y = y
}
