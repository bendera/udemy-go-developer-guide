package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	tr := triangle{base: 100, height: 50}
	sq := square{sideLength: 40}

	printArea(tr)
	printArea(sq)
}
